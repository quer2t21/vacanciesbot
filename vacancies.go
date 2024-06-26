package main

type Vacancy struct {
	name   string
	salary int
	url    string
	reqs   []int
}

var Vacancies []Vacancy = []Vacancy{
	{
		name:   "Юрисконсульт",
		salary: 70000,
		url:    "https://gauctr.ru/job/yuriskonsult-2/",
		reqs:   []int{2, 1, 1, 1},
	},
	{
		name:   "Экономист по договорной и претензионной работе",
		salary: 25000,
		url:    "https://gauctr.ru/job/ekonomist-po-dogovornoj-i-pretenzionnoj-rabote-ot-25-000-rub/",
		reqs:   []int{2, 1, 1, 1},
	},
	{
		name:   "Ведущий инспектор отдела трудовой миграции из стран с визовым порядком въезда",
		salary: 30000,
		url:    "https://gauctr.ru/job/vedushhij-inspektor-otdela-trudovoj-migratsii-iz-stran-s-vizovym-poryadkom-vezda/",
		reqs:   []int{2, 1, 1, 1},
	},
	{
		name:   "Ведущий инженер по эксплуатации зданий",
		salary: 35000,
		url:    "https://gauctr.ru/job/vedushhij-inzhener-po-ekspluatatsii-zdanij/",
		reqs:   []int{2, 1, 1, 1},
	},
	{
		name:   "Специалист по закупкам",
		salary: 35000,
		url:    "https://gauctr.ru/job/spetsialist-po-zakupkam/",
		reqs:   []int{2, 1, 1, 1},
	},
	{
		name:   "Cпециалист по закупкам",
		salary: 35000,
		url:    "https://gauctr.ru/job/spetsialist-po-zakupkam/",
		reqs:   []int{3, 1, 1, 1},
	},
	{
		name:   "Дворник, подсобный рабочий",
		salary: 40000,
		url:    "https://gauctr.ru/job/dvornik-podsobnyj-rabochij/",
		reqs:   []int{5, 0, 1, 1},
	},
	{
		name:   "Слесарь-сантехник 4 разряда",
		salary: 49000,
		url:    "https://gauctr.ru/job/slesar-santehnik-4-razryada/",
		reqs:   []int{5, 0, 1, 1},
	},
	{
		name:   "Техник-смотритель",
		salary: 67500,
		url:    "https://gauctr.ru/job/tehnik-smotritel/",
		reqs:   []int{5, 0, 1, 0},
	},
	{
		name:   "Маляр 4 разряда",
		salary: 49000,
		url:    "https://gauctr.ru/job/malyar-4-razryada/",
		reqs:   []int{5, 0, 1, 0},
	},
	{
		name:   "Санитар",
		salary: 35000,
		url:    "https://gauctr.ru/job/sanitar-ot-35-000-rublej/",
		reqs:   []int{6, 1, 1, 1},
	},
	{
		name:   "Инструктор по лечебной физкультуре",
		salary: 45000,
		url:    "https://gauctr.ru/job/instruktor-po-lechebnoj-fizkulture-ot-45-000-rublej/",
		reqs:   []int{6, 1, 1, 1},
	},
	{
		name:   "Медицинская сестра процедурной",
		salary: 45000,
		url:    "https://gauctr.ru/job/meditsinskaya-sestra-protsedurnoj-ot-45-000-rublej/",
		reqs:   []int{6, 1, 1, 1},
	},
	{
		name:   "Инструктор по адаптивной физической культуре",
		salary: 32000,
		url:    "https://gauctr.ru/job/instruktor-po-adaptivnoj-fizicheskoj-kulture-ot-32-000-do-60-000-rublej/",
		reqs:   []int{6, 1, 1, 0},
	},
	{
		name:   "Учитель английского языка",
		salary: 60000,
		url:    "https://gauctr.ru/job/uchitel-anglijskogo-yazyka-ot-60-000-rub/",
		reqs:   []int{7, 1, 1, 1},
	},
	{
		name:   "Педагог дополнительного образования",
		salary: 0,
		url:    "https://gauctr.ru/job/pedagog-dopolnitelnogo-obrazovaniya/",
		reqs:   []int{7, 1, 1, 1},
	},
	{
		name:   "Учитель химии",
		salary: 0,
		url:    "https://gauctr.ru/job/uchitel-himii/",
		reqs:   []int{7, 1, 1, 1},
	},
	{
		name:   "Наладчик станков и манипуляторов с программным управлением",
		salary: 120000,
		url:    "https://gauctr.ru/job/naladchik-stankov-i-manipulyatorov-s-programmnym-upravleniem-140-000-rub/",
		reqs:   []int{7, 1, 1, 1},
	},
	{
		name:   "Инженер-технолог производства радиоэлектронной аппаратуры",
		salary: 90000,
		url:    "https://gauctr.ru/job/inzhener-tehnolog-proizvodstva-radioelektronnoj-apparatury-ot-90-000-rublej-na-ruki-i-vyshe/",
		reqs:   []int{7, 1, 1, 1},
	},
	{
		name:   "Монажник РЭА",
		salary: 80000,
		url:    "https://gauctr.ru/job/montazhnik-rea-obemnyj-montazh-uzlov-blokov-priborov-ot-80-000-do-140-000-rublej-na-ruki/",
		reqs:   []int{7, 1, 1, 1},
	},
	{
		name:   "Слесарь-сборщик радиоэлектронной аппаратуры",
		salary: 75000,
		url:    "https://gauctr.ru/job/slesar-sborshhik-radioelektronnoj-apparatury-rea-ot-75-000-rublej/",
		reqs:   []int{7, 0, 1, 1},
	},
	{
		name:   "Монтажник РЭА",
		salary: 80000,
		url:    "https://gauctr.ru/job/montazhnik-rea-montazh-smd-pechatnyh-plat-ot-80-000-do-120-000-rublej-na-ruki/",
		reqs:   []int{7, 0, 0, 0},
	},
	{
		name:   "Водитель автомобильного грузового категория E",
		salary: 87000,
		url:    "https://gauctr.ru/job/voditel-avtomobilya-gruzovogo-kategoriya-e-zp-ot-87-000/",
		reqs:   []int{7, 0, 0, 1},
	},
	{
		name:   "Дробильщик",
		salary: 67000,
		url:    "https://gauctr.ru/job/drobilshhik-zp-ot-67-000/",
		reqs:   []int{7, 0, 0, 1},
	},
	{
		name:   "Термист вакуумных печей",
		salary: 80000,
		url:    "https://gauctr.ru/job/termist-vakuumnyh-pechej-ot-80-000-rublej/",
		reqs:   []int{7, 1, 0, 0},
	},
	{
		name:   "Токар-расточник",
		salary: 119500,
		url:    "https://gauctr.ru/job/tokar-rastochnik-6-razryada/",
		reqs:   []int{7, 0, 0, 1},
	},
	{
		name:   "Инженер-технолог производства радиоэлектронной аппаратуры",
		salary: 90000,
		url:    "https://gauctr.ru/job/inzhener-tehnolog-proizvodstva-radioelektronnoj-apparatury-ot-90-000-rublej-na-ruki-i-vyshe/",
		reqs:   []int{7, 0, 1, 0},
	},
}
