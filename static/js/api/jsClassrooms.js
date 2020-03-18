// addClassroom = () =>{
// 	let number = document.querySelector('.numberCab').value
// 	let	placeQuantity = document.querySelector('.placeQuantity').value
// 	let	 building = document.querySelectorAll('.building-radio');
// 	let radioValue
// 	for(let i=0; i<building.length; i++ ){
// 		if (building[i].checked) {
// 			radioValue = building[i].value

// 		}
// 	}
// 	let isComputer = document.querySelector('.computerclass').checked

// 	var myHeaders = new Headers();
// 	myHeaders.append("Content-Type", "application/json");

// 	var raw = JSON.stringify({"PlaceQuantity":Number(placeQuantity),"IsComputer":isComputer,"IDBuilding":Number(radioValue),"Name": number});

// 		var requestOptions = {
// 		  method: 'POST',
// 		  headers: myHeaders,
// 		  body: raw,
// 		  redirect: 'follow'
// 		};
// 	url = 'http://schedule.tomtit.tomsk.ru/api/classroom/'
// 	fetch(url, requestOptions)
// 	  .then(response => response.text())
// 	  .then(result => console.log(result))
// 	  .catch(error => console.log('error', error));
	
// 	alert("Добавлено")
// }
