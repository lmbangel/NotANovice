package handlers

//func HandleGetQuestions(w http.ResponseWriter, r *http.Request) {
//	var qs db.Question
//	conn, err := sql.Open("sqlite", "./quiz.db")
//	if err != nil {
//		panic(err)
//	}

//	q := db.New(conn)
//	qs, err = q.GetQuestion(context.Background(), 1)
//	if err != nil {
//		if err.Error() == "sql: no rows in result set" {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusNotFound)
//			json.NewEncoder(w).Encode(map[string]string{
//				"message": "Question not found.",
//			})
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(err.Error())
//		return
//	}

//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(qs)
//}

//func HandleAnswersToQuestions(w http.ResponseWriter, r *http.Request) {
//	var answer db.Attempt

//	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
//		log.Fatalf("Error getting user answer: %s", err)
//	}

//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(answer)
//}
