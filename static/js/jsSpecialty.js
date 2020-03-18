addSpecialty = ()=>{
	let code = document.querySelector('.code').value
	let	name = document.querySelector('.name').value
	let	duration = document.querySelectorAll('.duration-radio');
	let radioValue
	for(let i=0; i<building.length; i++ ){
		if (duration[i].checked) {
			radioValue = duration[i].value

		}
	}

	var myHeaders = new Headers();
	myHeaders.append("Content-Type", "application/json");

	var raw = JSON.stringify({"Code":code,"Name":name,"IDDuration":Number(radioValue)});

		var requestOptions = {
		  method: 'POST',
		  headers: myHeaders,
		  body: raw,
		  redirect: 'follow'
		};
	url = 'http://schedule.tomtit.tomsk.ru/api/specialty/'
	fetch(url, requestOptions)
	  .then(response => response.text())
	  .then(result => console.log(result))
	  .catch(error => console.log('error', error));
	
	alert("Добавлено")
}