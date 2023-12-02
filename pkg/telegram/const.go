package telegram

const (
	unknownCommand = "Я не знаю такой команды!"
	start          = "start"
	startMessage   = "Добро пожаловать!\n" +
		"Введите номер вашей группы"
	help        = "help"
	helpMessage = "Я помогу тебе узнать расписание на день/неделю\n" +
		"Если группа не находится, то введи вместо английской 'C' русскую\n" +
		"Пока только для ИиВТ"
	groupExample         = "[СБC]\\d{2}-\\d{3}-\\d"
	fileIiVT             = "IiVT.xlsx"
	theGroupHasBeenFound = "Группа найдена"
	theGroupWasNotfound  = "Группа не найдена\n" +
		"Попробуйте вместо английской 'C' русскую"
	weekOverUnder = "Неделя над/под"
)
