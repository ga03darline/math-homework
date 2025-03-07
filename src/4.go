func getRandomNumber() int {
	return rand.Intn(10) + 1
}

func getSumOfNumbers() int {
	return getRandomNumber() + getRandomNumber()
}

func main() {
	fmt.Println(getSumOfNumbers())
}