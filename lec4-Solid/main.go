package main

type List interface {
	Add(el int)
	Remove()
	Size() int
	Get(idx int) int
	Set(newEl int)
}

type ArrayList struct {
	arr []int
}

func (array *ArrayList) Add(el int) {
	array.arr = append(array.a, el)
}
func (array *ArrayList) Size() int {
	return len(array.arr)
}
func (array *ArrayList) Get(idx int) int {
	//if array.Size() > idx {
	//	return array.arr[idx]
	//}
	//return
	return array.arr[idx]
}
func (array *ArrayList) Set(idx int,newel int)  {
	array.arr[idx] = newel
}
type Vector struct {
	arr []int
}
func (v *Vector) Add(el int) {
	v.arr = append(v.arr, el)
}
func (v *Vector) Size() int {
	return len(v.arr)
}
func (v *Vector) Get(idx int) int {
	return v.arr[idx]
}
func (v *Vector) Set(idx int, newel int)  {
	v.arr[idx] = newel
}
func main() {

}
