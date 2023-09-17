package game

import (
	"fmt"
	"golang_beginning/lesson_03/env"
)

func Game() {
	fmt.Println("Вітаємо у грі!")
	fmt.Println("Ви опинилися в лісі. Що ви будете робити?")
	fmt.Println("1. Піти вглиб лісу")
	fmt.Println("2. Повернутися назад")
	fmt.Println("3. Знайти місце для ночівлі")

	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Ви зустріли ведмедя. Що ви будете робити?")
		fmt.Println("1. Захиститися")
		fmt.Println("2. Втікти")
		fmt.Println("3. Спробувати зв'язатися з ним")

		var choice string
		fmt.Scanln(&choice)

		if choice == "1" {
			fmt.Println("Ви використали свої навички боротьби з ведмедем і вижили!")
		} else if choice == "2" {
			fmt.Println("Ви втекли від ведмедя, але загубилися в лісі.")
		} else if choice == "3" {
			fmt.Println("Ви змогли зв'язатися з ведмедем і він виявився дружелюбним.")
		} else {
			fmt.Println("Ви нічого не зробили і ведмідь вас з'їв.")
			goto exit
		}
	case 2:
		fmt.Println("Ви повернулися до початку лісу і зустрілися з лісовими рейнджерами.")
		fmt.Println("1. Попросити допомоги")
		fmt.Println("2. Спробувати втекти")
		fmt.Println("3. Спробувати переконати їх, що ви не потребуєте допомоги")

		var choice string
		fmt.Scanln(&choice)

		if choice == "1" {
			fmt.Println("Лісові рейнджери допомогли вам вийти з лісу.")
		} else if choice == "2" {
			fmt.Println("Ви спробували втекти, але лісові рейнджери зловили вас.")
		} else if choice == "3" {
			fmt.Println("Лісові рейнджери повірили вам і пішли далі.")
		} else {
			goto exit
		}
	case 3:
		fmt.Println("Ви знайшли місце для ночівлі, але вас затягнуло в сон.")
		fmt.Println("1. Прокинутися")
		fmt.Println("2. Продовжити спати")

		var choice string
		fmt.Scanln(&choice)

		if choice == "1" {
			fmt.Println("Ви прокинулися від шуму і знайшли водопад.")
			fmt.Print("Зайти у водопад? (y/n): ")
			fmt.Scanln(&choice)
			if choice == "y" {
				fmt.Println("Ви зайшли у водопад і знайшли сейф з паролем.")
				fmt.Println("Введіть однозначний пароль: ")
				var password string

				for password != secret.SecretPassword {
					fmt.Scanln(&password)
					if password != secret.SecretPassword {
						fmt.Println("Неправильний пароль. Спробуйте ще раз.")
					}
				}
				fmt.Println("Ви відкрили сейф і знайшли там скарби!")
				goto exit
			} else {
				fmt.Println("Ви не зайшли у водопад і вас з'їли комарі.")
			}

		} else if choice == "2" {
			fmt.Println("Ви продовжили спати і проспали весь день.")
		} else {
			goto exit
		}
	default:
		goto exit
	}
exit:
	fmt.Println("Гра завершена. Вас з'їли комарі.")
}
