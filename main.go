package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type requirements struct {
	chatID telego.ChatID
	reqs   []int
	salary int
}

const (
	IT                         = 1
	AdminWork                  = 2
	Accounting                 = 3
	Delivery                   = 4
	HousingAndCommunalServices = 5
	Medicine                   = 6
	Education                  = 7
	Manufacture                = 8
	DistantWork                = 1
	SNG                        = 1
	Guide                      = 1
)

func main() {
	bot, err := telego.NewBot(mustToken())
	if err != nil {
		panic(err)
	}
	fmt.Print("Bot is running now")

	updates, err := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()
	if err != nil {
		panic(err)
	}
	var reqs requirements
	reqs.reqs = []int{0, 0, 0, 0}
	for update := range updates {
		var chatID telego.ChatID
		var msgID int
		if update.Message != nil {
			chatID = update.Message.Chat.ChatID()
			fmt.Print(chatID)
		}

		if update.CallbackQuery != nil {
			query := update.CallbackQuery
			chat := query.Message.GetChat()
			chatID = chat.ChatID()
			msgID = update.CallbackQuery.Message.GetMessageID()
		}
		reqs.chatID = chatID
		if (update.Message != nil && update.Message.Text == "/start") || (update.CallbackQuery != nil && update.CallbackQuery.Data == "1") {
			reqs.reqs = []int{0, 0, 0, 0}
			startMessage := "Привет! Я бот по поиску ваканский, заполните небольшую анкету из 5 вопросов и я подберу подходящие вакансии\nВыберите специальность:"
			button1 := tu.InlineKeyboardButton("IT-Сфера").WithCallbackData("IT")
			button2 := tu.InlineKeyboardButton("Административная работа").WithCallbackData("AdminWork")

			button3 := tu.InlineKeyboardButton("Бухгалтерия").WithCallbackData("Accounting")
			button4 := tu.InlineKeyboardButton("Доставка").WithCallbackData("Delivery")

			button5 := tu.InlineKeyboardButton("ЖКХ").WithCallbackData("HousingAndCommunalServices")
			button6 := tu.InlineKeyboardButton("Медицина").WithCallbackData("Medicine")

			button7 := tu.InlineKeyboardButton("Образование").WithCallbackData("Education")
			button8 := tu.InlineKeyboardButton("Производство").WithCallbackData("Manufacture")

			row := tu.InlineKeyboardRow(button1, button2)
			row2 := tu.InlineKeyboardRow(button3, button4)
			row3 := tu.InlineKeyboardRow(button5, button6)
			row4 := tu.InlineKeyboardRow(button7, button8)
			replyMarkup := tu.InlineKeyboard(row, row2, row3, row4)

			msgParams := telego.SendMessageParams{
				ChatID:      chatID,
				Text:        startMessage,
				ReplyMarkup: replyMarkup,
			}
			_, err := bot.SendMessage(&msgParams)
			if err != nil {
				log.Fatal(err)
			}

		}

		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "IT":
				reqs.reqs[0] = IT
			case "AdminWork":
				reqs.reqs[0] = AdminWork
			case "Accounting":
				reqs.reqs[0] = Accounting
			case "Delivery":
				reqs.reqs[0] = Delivery
			case "HousingAndCommunalServices":
				reqs.reqs[0] = HousingAndCommunalServices
			case "Medicine":
				reqs.reqs[0] = Medicine
			case "Education":
				reqs.reqs[0] = Education
			case "Manufacture":
				reqs.reqs[0] = Manufacture
			case "DistantWork":
				reqs.reqs[1] = DistantWork
			case "NoDistantWork":
				reqs.reqs[1] = 0
			case "SNG":
				reqs.reqs[2] = SNG
			case "NoSNG":
				reqs.reqs[2] = 0
			case "Guide":
				reqs.reqs[3] = Guide
			case "NoGuide":
				reqs.reqs[3] = 0
			}
		}

		if update.CallbackQuery != nil && (update.CallbackQuery.Data == "IT" ||
			update.CallbackQuery.Data == "AdminWork" ||
			update.CallbackQuery.Data == "Accounting" ||
			update.CallbackQuery.Data == "Delivery" ||
			update.CallbackQuery.Data == "HousingAndCommunalServices" ||
			update.CallbackQuery.Data == "Medicine" ||
			update.CallbackQuery.Data == "Education" ||
			update.CallbackQuery.Data == "Manufacture") {
			message := "Возможность удаленной работы?"
			button1 := tu.InlineKeyboardButton("Да").WithCallbackData("DistantWork")
			button2 := tu.InlineKeyboardButton("Необязательно").WithCallbackData("NoDistantWork")
			buttonBack := tu.InlineKeyboardButton("Заново").WithCallbackData("1")
			row := tu.InlineKeyboardRow(button1, button2)
			rowBack := tu.InlineKeyboardRow(buttonBack)
			replyMarkup := tu.InlineKeyboard(row, rowBack)

			msgParams := telego.EditMessageTextParams{
				ChatID:      chatID,
				MessageID:   msgID,
				Text:        message,
				ReplyMarkup: replyMarkup,
			}
			_, err := bot.EditMessageText(&msgParams)
			if err != nil {
				log.Fatal(err)
			}
		}

		if update.CallbackQuery != nil && (update.CallbackQuery.Data == "DistantWork" ||
			update.CallbackQuery.Data == "NoDistantWork") {
			message := "Показывать вакансии только для граждан СНГ?"
			button1 := tu.InlineKeyboardButton("Да").WithCallbackData("SNG")
			button2 := tu.InlineKeyboardButton("Нет").WithCallbackData("NoSNG")
			buttonBack := tu.InlineKeyboardButton("Заново").WithCallbackData("1")
			row := tu.InlineKeyboardRow(button1, button2)
			rowBack := tu.InlineKeyboardRow(buttonBack)
			replyMarkup := tu.InlineKeyboard(row, rowBack)

			msgParams := telego.EditMessageTextParams{
				ChatID:      chatID,
				MessageID:   msgID,
				Text:        message,
				ReplyMarkup: replyMarkup,
			}
			_, err := bot.EditMessageText(&msgParams)
			if err != nil {
				log.Fatal(err)
			}
		}

		if update.CallbackQuery != nil && (update.CallbackQuery.Data == "SNG" ||
			update.CallbackQuery.Data == "NoSNG") {
			message := "Вас интересует работа с обучением?"
			button1 := tu.InlineKeyboardButton("Да").WithCallbackData("Guide")
			button2 := tu.InlineKeyboardButton("Нет").WithCallbackData("NoGuide")
			buttonBack := tu.InlineKeyboardButton("Заново").WithCallbackData("1")
			row := tu.InlineKeyboardRow(button1, button2)
			rowBack := tu.InlineKeyboardRow(buttonBack)
			replyMarkup := tu.InlineKeyboard(row, rowBack)

			msgParams := telego.EditMessageTextParams{
				ChatID:      chatID,
				MessageID:   msgID,
				Text:        message,
				ReplyMarkup: replyMarkup,
			}
			_, err := bot.EditMessageText(&msgParams)
			if err != nil {
				log.Fatal(err)
			}
		}

		if update.CallbackQuery != nil && (update.CallbackQuery.Data == "Guide" ||
			update.CallbackQuery.Data == "NoGuide") {
			message := "Желаемая зарплата: "
			button1 := tu.InlineKeyboardButton("25000").WithCallbackData("25000")
			button2 := tu.InlineKeyboardButton("50000").WithCallbackData("50000")
			button3 := tu.InlineKeyboardButton("100000").WithCallbackData("100000")
			button4 := tu.InlineKeyboardButton("150000").WithCallbackData("150000")
			button5 := tu.InlineKeyboardButton("200000").WithCallbackData("200000")
			button6 := tu.InlineKeyboardButton("Неважно").WithCallbackData("0")
			buttonBack := tu.InlineKeyboardButton("Заново").WithCallbackData("1")
			row := tu.InlineKeyboardRow(button1, button2)
			row2 := tu.InlineKeyboardRow(button3, button4)
			row3 := tu.InlineKeyboardRow(button5, button6)
			rowBack := tu.InlineKeyboardRow(buttonBack)
			replyMarkup := tu.InlineKeyboard(row, row2, row3, rowBack)

			msgParams := telego.EditMessageTextParams{
				ChatID:      chatID,
				MessageID:   msgID,
				Text:        message,
				ReplyMarkup: replyMarkup,
			}
			_, err := bot.EditMessageText(&msgParams)
			if err != nil {
				log.Fatal(err)
			}
		}

		if update.CallbackQuery != nil && (update.CallbackQuery.Data == "25000" ||
			update.CallbackQuery.Data == "50000" ||
			update.CallbackQuery.Data == "100000" ||
			update.CallbackQuery.Data == "150000" ||
			update.CallbackQuery.Data == "200000" ||
			update.CallbackQuery.Data == "0") {
			salary, _ := strconv.Atoi(update.CallbackQuery.Data)
			foundVacs := findVacancy(reqs, salary)
			if len(foundVacs) == 0 {
				vacmsgParams := telego.SendMessageParams{
					ChatID: reqs.chatID,
					Text:   "Подходящих вакансий не найдено.",
				}
				_, err := bot.SendMessage(&vacmsgParams)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				mParams := telego.SendMessageParams{
					ChatID: reqs.chatID,
					Text:   "Подобранные вакансии:",
				}
				_, err := bot.SendMessage(&mParams)
				if err != nil {
					log.Fatal(err)
				}
			}
			for _, vac := range foundVacs {
				vacmsgParams := telego.SendMessageParams{
					ChatID: reqs.chatID,
					Text:   vac,
				}
				_, err := bot.SendMessage(&vacmsgParams)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		fmt.Print(reqs.reqs)

	}

}

func findVacancy(reqs requirements, salary int) []string {
	var vacUrls []string
	for _, vac := range Vacancies {
		if vac.reqs[0] == reqs.reqs[0] &&
			(vac.reqs[1] == reqs.reqs[1] || reqs.reqs[1] == 0) &&
			vac.reqs[2] == reqs.reqs[2] &&
			vac.reqs[3] == reqs.reqs[3] &&
			vac.salary >= salary {
			vacUrls = append(vacUrls, vac.name+"\nЗарплата: "+strconv.Itoa(vac.salary)+"\n"+vac.url)
		}
	}
	return vacUrls
}
func mustToken() string {
	bot_token := os.Getenv("TG_BOT_TOKEN") // токен хранится в переменных среды
	token := flag.String("token", bot_token, "tg-bot-token")

	flag.Parse()

	if *token == "" {
		panic("token is empty")
	}

	return *token
}
