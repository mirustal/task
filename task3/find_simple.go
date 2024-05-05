package findSimple

func CheckSimple(number int)  bool{
    isSimple := true
    if number == 0 || number == 1 {
        return false
    } else {
        for i := 2;
        i <= number / 2;
        i++ {
            if number % i == 0 {
                isSimple = false
                break
            }
        }
        
    }
	if isSimple == true {
		return true
	}
	return false
}

func CheckListSimple(min int, max int) []int{ 
    var simpleNums []int
    for i := min; i <= max; i++ {
        if CheckSimple(i) {
            simpleNums = append(simpleNums, i)
        }
    }

    return simpleNums
}

