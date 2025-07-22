CRUD с интерфейсами
Реализация CRUD с интерфейсами. 

Запуск: 
go run main.go

Пример:
service.Create(User{1, "Иван"})
service.GetByID(1)
service.UpdateUser(User{1, "Антон"})
service.DeleteUser(User{1, "Антон"})
