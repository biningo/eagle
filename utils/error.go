package utils

/**
*@Author icepan
*@Date 7/20/21 16:15
*@Describe
**/

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
