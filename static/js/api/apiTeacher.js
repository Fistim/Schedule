addTeacher = () =>{
    let surname = document.querySelector('.surname').value
    let name = document.querySelector('.name').value
    let patronymic = document.querySelector('.patronymic').value
	    var select = document.getElementById("selectClassroom");
	    var value = select.value;

	var myHeaders = new Headers();
	myHeaders.append("Content-Type", "application/json");

	var raw = {"Surname":surname,"Name":name,"IDCLassroom":Number(value),"Patronymic": patronymic};
	requestFunc(url="/teacher/", method="POST", data=raw)
	alert("Добавлено")
}