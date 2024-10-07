package utils

import (
	"backend/src/internal/model"
	"fmt"
)

//func PrintHeader(val any) {
//	var str string
//	t := reflect.TypeOf(val)
//
//	for i := 0; i < t.NumField(); i++ {
//		str = fmt.Sprintf("%s | %s", str, t.Field(i).Name)
//	}
//	fmt.Println(str)
//}

func PrintReserves(reserves []*model.Reserve) {
	//var str string
	for i := 0; i < len(reserves); i++ {
		fmt.Printf("%d. Комната: %s\n   Продюсер: %s\n   Инструменталист: %s\n   Время начала: %s\n   Время конца: %s\n",
			i+1,
			reserves[i].RoomId,
			reserves[i].ProducerId,
			reserves[i].InstrumentalistId,
			reserves[i].TimeInterval.StartTime.String(),
			reserves[i].TimeInterval.EndTime.String(),
		)
	}

	//fmt.Printf("%s\n", str)
}

func PrintStudios(studios []*model.Studio) {
	var str string
	i := 1
	for _, studio := range studios {
		str += fmt.Sprintf("%d. %s\n", i, studio.Name)
		i++
	}

	fmt.Printf("%s\n", str)
}

func PrintRooms(rooms []*model.Room) {
	var str string
	i := 1
	for _, room := range rooms {
		str += fmt.Sprintf("%d. %s\n", i, room.Name)
		i++
	}

	fmt.Printf("%s\n", str)
}

func PrintProducers(producers []*model.Producer) {
	var str string
	i := 1
	for _, producer := range producers {
		str += fmt.Sprintf("%d. %s\n", i, producer.Name)
		i++
	}

	fmt.Printf("%s\n", str)
}

func PrintInstrumentalists(instrumentalists []*model.Instrumentalist) {
	var str string
	i := 1
	for _, instrumentalist := range instrumentalists {
		str += fmt.Sprintf("%d. %s\n", i, instrumentalist.Name)
		i++
	}

	fmt.Printf("%s\n", str)
}

func PrintEquipments(equipments []*model.Equipment) {
	var str string
	i := 1
	for _, equipment := range equipments {
		str += fmt.Sprintf("%d. %s\n", i, equipment.Name)
		i++
	}

	fmt.Printf("%s\n", str)
}

//
//func PrintCollection[T any](collName string, collection []*T) {
//	fmt.Println(collName)
//
//	for i, val := range collection {
//		if i == 0 {
//			PrintHeader(*val)
//		}
//
//		if reflect.TypeOf(val).Kind() == reflect.Ptr {
//			PrintStruct(reflect.ValueOf(val).Elem().Interface())
//		} else {
//			PrintStruct(val)
//		}
//	}
//}
//
//func PrintActivityField(field *domain.ActivityField) {
//	fmt.Printf("%s | %s | %s | %f\n", field.ID, field.Name, field.Description, field.Cost)
//}
//
//func PrintActivityFields(fields []*domain.ActivityField) {
//	fmt.Println("Сферы деятельности:")
//	for _, field := range fields {
//		PrintActivityField(field)
//	}
//}
//
//func PrintPaginatedCollectionArgs[T any](collectionName string, fn func(uuid.UUID, int) ([]*T, error), id uuid.UUID) (err error) {
//	page := 1
//
//	for {
//		tmp, err := fn(id, page)
//		if err != nil {
//			return fmt.Errorf("получение пагинированных данных: %w", err)
//		}
//
//		PrintCollection(collectionName, tmp)
//
//		fmt.Printf("1. Предыдущая страница.\n2. Следующая страница.\n0. Назад.\n\nВыберите действие: ")
//		var option int
//		_, err = fmt.Scanf("%d", &option)
//		if err != nil {
//			return fmt.Errorf("ошибка ввода следующего действия: %w", err)
//		}
//
//		switch option {
//		case 1:
//			if page > 1 {
//				page--
//			}
//		case 2:
//			if len(tmp) == config.PageSize {
//				page++
//			}
//		case 0:
//			return nil
//		}
//	}
//}
//
//func PrintPaginatedCollection[T any](collectionName string, fn func(int) ([]*T, error)) (err error) {
//	page := 1
//
//	for {
//		tmp, err := fn(page)
//		if err != nil {
//			return fmt.Errorf("получение пагинированных данных: %w", err)
//		}
//
//		PrintCollection(collectionName, tmp)
//
//		fmt.Printf("1. Предыдущая страница.\n2. Следующая страница.\n0. Назад.\n\nВыберите действие: ")
//		var option int
//		_, err = fmt.Scanf("%d", &option)
//		if err != nil {
//			return fmt.Errorf("ошибка ввода следующего действия: %w", err)
//		}
//
//		switch option {
//		case 1:
//			if page > 1 {
//				page--
//			}
//		case 2:
//			if len(tmp) == config.PageSize {
//				page++
//			}
//		case 0:
//			return nil
//		}
//	}
//}

//func PrintYearCollection[T any](collectionName string, fn func(uuid.UUID, *domain.Period) (*T, error), id uuid.UUID) (err error) {
//	curYear := time.Now().Year()
//	curPeriod := &domain.Period{
//		StartYear:    curYear,
//		EndYear:      curYear,
//		StartQuarter: 1,
//		EndQuarter:   4,
//	}
//
//	for {
//		actYear := time.Now()
//		tmp, err := fn(id, curPeriod)
//		if err != nil {
//			return fmt.Errorf("получение периодизированных данных: %w", err)
//		}
//
//		PrintCollection(collectionName, tmp)
//
//		fmt.Printf("1. Предыдущий год.\n2. Следующий год.\n0. Назад.\n\nВыберите действие: ")
//		var option int
//		_, err = fmt.Scanf("%d", &option)
//		if err != nil {
//			return fmt.Errorf("ошибка ввода следующего действия: %w", err)
//		}
//
//		switch option {
//		case 1:
//			if actYear.Year() <= curYear {
//				actYear.AddDate(-1, 0, 0)
//			}
//		case 2:
//			if actYear.Year() >= curYear {
//				actYear.AddDate(1, 0, 0)
//			}
//		case 0:
//			return nil
//		}
//	}
//}
