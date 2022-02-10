type printer interface(){
	//behaviour is of print
	print()  //just a method name without code
}
type list[]printer
func (l list) print(){
	if len(l)===0{
		fmt.Println("sorry")
		return
	}
for _,it:=range l{
	it.print()
}
}