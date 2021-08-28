package entity

func RemovePlayerRoleFromSlice(s []IRole, index int) []IRole {
    ret := make([]IRole, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

//https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang