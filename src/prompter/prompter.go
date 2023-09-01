package prompter

import "fmt"

func SelectWithCallBack(prompt string, options []string, functions []func() error) error {
	if len(options) != len(functions) {
		return fmt.Errorf("number of function and option need to be the same: %d opt - %d fn", len(options), len(functions))
	}
	fmt.Printf("%s\n", prompt)
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Printf("Enter a number: ")
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		return err
	}
	if input < 1 || input > len(options) {
		return fmt.Errorf("invalid input: %d", input)
	}
	err = functions[input-1]()
	return err
}

func Select(prompt string, options []string) int {
	fmt.Printf("%s\n", prompt)
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Printf("Enter a number: ")
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		return -1
	}
	if input < 1 || input > len(options) {
		return 0
	}
	return input
}

func Confirm(prompt string) (bool, error) {
	fmt.Printf("%s [y/n]: ", prompt)
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return false, err
	}
	if input == "y" {
		return true, nil
	} else if input == "n" {
		return false, nil
	}
	return false, fmt.Errorf("invalid input: %s", input)
}

func Input(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return "", err
	}
	return input, nil
}

func InputInt(prompt string) (int, error) {
	fmt.Printf("%s: ", prompt)
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		return 0, err
	}
	return input, nil
}

func InputFloat(prompt string) (float64, error) {
	fmt.Printf("%s: ", prompt)
	var input float64
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		return 0, err
	}
	return input, nil
}
