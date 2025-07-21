package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	Create(user User) error
	GetByID(ID int) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(user User) error
}

type MemoryRepo struct {
	data map[int]User
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{data: make(map[int]User)}
}

func (m *MemoryRepo) Create(user User) error {
	m.data[user.ID] = user
	return nil
}

func (m *MemoryRepo) GetByID(ID int) (User, error) {
	user, ok := m.data[ID]
	if !ok {
		return User{}, fmt.Errorf("Пользователь не найден")
	}
	return user, nil
}

func (m *MemoryRepo) UpdateUser(user User) (User, error) {
	_, ok := m.data[user.ID]
	if !ok {
		return User{}, fmt.Errorf("Пользователь не найден")
	}
	m.data[user.ID] = user
	return user, nil
}

func (m *MemoryRepo) DeleteUser(user User) error {
	_, ok := m.data[user.ID]
	if !ok {
		return fmt.Errorf("Пользователь не найден")
	}
	delete(m.data, user.ID)
	return nil
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user User) (User, error) {
	err := s.repo.Create(user)
	if err != nil {
		return User{}, fmt.Errorf("Не удалось создать пользователя %s, %v", user.Name, err)
	}
	fmt.Printf("Пользователь создан: %s\n", user.Name)
	return user, nil
}

func (s *UserService) GetByID(ID int) (User, error) {
	user, err := s.repo.GetByID(ID)
	if err != nil {
		return User{}, fmt.Errorf("Не удалось найти пользователя %d, %v", ID, err)
	}
	fmt.Printf("Пользователь %s\n", user.Name)
	return user, nil
}

func (s *UserService) UpdateUser(user User) (User, error) {
	user, err := s.repo.UpdateUser(user)
	if err != nil {
		return User{}, fmt.Errorf("Невозможно обновить пользователя %s, %v", user, err)
	}
	fmt.Printf("Пользователь обновлен: %s\n", user.Name)
	return user, nil
}

func (s *UserService) DeleteUser(user User) error {
	err := s.repo.DeleteUser(user)
	if err != nil {
		return fmt.Errorf("Не удалось удалить пользователя %s, %v", user.Name, err)
	}
	fmt.Printf("Пользователь %s удален\n", user.Name)
	return nil
}

func main() {
	repo := NewMemoryRepo()
	service := NewUserService(repo)

	_, err := service.Create(User{1, "Иван"})
	if err != nil {
		fmt.Println("Не удалось создать пользователя")
		return
	}
	_, err = service.GetByID(1)
	if err != nil {
		fmt.Println("Не удалось опознать ID")
		return
	}
	_, err = service.UpdateUser(User{1, "Антон"})
	if err != nil {
		fmt.Println("Не удалось обновить пользователя")
		return
	}

	err = service.DeleteUser(User{1, "Антон"})
	if err != nil {
		fmt.Println("Не удалось удалить пользователя")
		return
	}
}
