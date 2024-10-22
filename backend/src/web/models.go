package web

//func ToReserveModel(reserveRequest *CreateReserveRequest) *model.Reserve {
//	startTimeString := strings.TrimSpace(reserveRequest.Date) + " " +
//		strings.TrimSpace(reserveRequest.StartHour) + ":00:00"
//
//	endTimeString := strings.TrimSpace(reserveRequest.Date) + " " +
//		strings.TrimSpace(reserveRequest.EndHour) + ":00:00"
//
//	//_, err = time_parser.StringToDate(startTimeString)
//	startTime, err := time_parser.StringToDate(startTimeString)
//	//if err != nil {
//	//	errorTextView.SetText(err.Error())
//	//	tui.ErrorForm(errorForm, pages, errorTextView, CreateValidatePage)
//	//	pages.SwitchToPage(ErrorPage)
//	//}
//	//_, err = time_parser.StringToDate(endTimeString)
//	endTime, err := time_parser.StringToDate(endTimeString)
//	//if err != nil {
//	//	errorTextView.SetText(err.Error())
//	//	tui.ErrorForm(errorForm, pages, errorTextView, CreateValidatePage)
//	//	pages.SwitchToPage(ErrorPage)
//	//}
//
//	userInterval := serviceImpl.NewTimeInterval(startTime, endTime)
//
//	userIdInt, err := strconv.Atoi(reserveRequest.UserId)
//	roomIdInt, err :=
//	//if err != nil {
//	//	//web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: преобразование id пользователя к int: %w", prompt, err).Error(), http.StatusBadRequest)
//	//	return
//	//}
//
//	return model.Reserve{
//		UserId: int64(userIdInt),
//		RoomId:
//	}
//}
