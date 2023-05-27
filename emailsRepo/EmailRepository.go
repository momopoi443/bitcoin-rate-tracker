package emailsRepo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Add(email string) error {
	file, err := os.OpenFile("SubscribedEmails.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Додавання запису до файлу
	_, err = fmt.Fprintf(file, "%s\n", email)
	if err != nil {
		return err
	}

	return nil
}

func Exist(email string) (bool, error) {
	file, err := os.Open("SubscribedEmails.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == email {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func ListAll() ([]string, error) {
	file, err := os.Open("SubscribedEmails.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var emails []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emails = append(emails, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return emails, nil
}
