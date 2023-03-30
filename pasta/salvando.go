// You can edit this code!
// Click here and start typing.
import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	sum := func(a, b int) int { return a + b }(3, 4)
	fmt.Println(sum)

	var salario float32
	var salarioMaisBonus float32

	salario = 1.2

	salarioMaisBonus = salario

	if salario < 50 {
		salarioMaisBonus = (salario + 200)
	}

	fmt.Println(salarioMaisBonus)

	fmt.Println("======================================================== dividindo exercicio")

	texto := "palavra"
	fmt.Println("Quantidade:", len(texto))

	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "r" {
			break
		}
		fmt.Println(string(texto[i]))
	}

	fmt.Println("======================================================== dividindo exercicio")

	for numBase := 1; numBase <= 10; numBase++ {
		for multiplicador := 1; multiplicador <= 10; multiplicador++ {
			fmt.Println(numBase, "x", multiplicador, "=", numBase*multiplicador)
		}
	}

	fmt.Println("======================================================== dividindo exercicio")

	list := []int{1, 2, 3, 4, 5}
	fmt.Println(list)
	list = append(list, 10)
	fmt.Println(list)

	list1 := make([]int, 1)
	fmt.Println(list1)

	fmt.Println("======================================================== dividindo exercicio")

	list2 := make([]int, 1)
	list2[0] = 5
	list2 = append(list2, 3)
	list2 = append(list2, 5)

	somaTotal := 0

	for i := 0; i < len(list2); i++ {
		somaTotal += list2[i]
	}

	fmt.Println("soma e", somaTotal/len(list2))

	fmt.Println("======================================================== dividindo exercicio")
	initial := []int{4, 3, 2, 1, 5, 10}
	initialEmpty := make([]int, 0)
	for i := 0; i < len(initial); i++ {
		if initial[i] < 5 {
			initialEmpty = append(initialEmpty, initial[i])
			fmt.Println("Posicao[", i, "]", initialEmpty)
		}
	}
	fmt.Println(initialEmpty)

	fmt.Println("======================================================== dividindo exercicio")

	listTotal := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	listFinal := listTotal[:3]
	listFinal1 := listTotal[4:]

	fmt.Println(listFinal)
	fmt.Println(listFinal1)

}
