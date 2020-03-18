
const requestFunc = async(url, method = "GET", data = null, token = null) => {
	apihost = 'http://schedule.tomtit.tomsk.ru/api'
	method = method.toLocaleUpperCase()
	let fullurl = `${apihost}${url}`;
	let options = {
	method: method,
	headers: {
	"Content-Type": "application/json",
	"Authorization": `Bearer ${token}`,
	},
};

switch(method) {
	case "PUT":
	delete options.headers["Content-Type"];
	options.body = data;
	break;
	case "POST": case "PATCH": case "DELETE":
	options.body = JSON.stringify(data);
	break;
}

const res = await fetch(fullurl, options);
return await res.json();
};
var app = new Vue({
        el: '#app',
        delimiters: ['${', '}'],
        data: {
            classrooms:[], 
            buildings: [],
            selectedRow: {
                "Code": "",
                "Name": "",
                "IDDuration": 0,
                "ID": ""
            },
         },
        methods: {
            showClassrooms() {
                console.log(this.classrooms)
            },
            selectRow(row, index){
                let rows = document.querySelectorAll('.clickHover')
                console.log(index)
                this.selectedRow = row;
                
                // for(let i = 0; i < row.length; i++) {
                //     rows[i].style.backgroundColor = 'red'
                // }
                console.log(rows)
                console.log(rows[index].style.backgroundColor)
                // empty string
                if(rows[index].style.backgroundColor != 'red'){
                    rows[index].style.backgroundColor = 'red'
                }
                else{
                    rows[index].style.backgroundColor = 'white'
                    this.clearSelectedRow()
                }   
                
                console.log(row)
            },
            clearSelectedRow(){
                this.selectedRow =   {
                "Code": "",
                "Name": "",
                "IDDuration": 0
                }
            },
        	async getClassrooms(){
        		this.classrooms = await requestFunc("/classroom", "GET")
        	}, 
            async getBuildings() {
                const buildings = await requestFunc("/building/", "GET")
                    .then(response => {
                        console.log(response)
                        this.buildings = response
                    })
            },
            
        }, 
        mounted() {
        	this.getClassrooms()
            this.getBuildings()
        }
    });

addClassroom = () =>{
    let number = document.querySelector('.numberCab').value
    let placeQuantity = document.querySelector('.placeQuantity').value
    let building = document.querySelectorAll('.building-radio');
    let radioValue
    for(let i=0; i<building.length; i++ ){
        if (building[i].checked) {
            radioValue = building[i].value

        }
    }
    let isComputer = document.querySelector('.computerclass').checked

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var raw = {"PlaceQuantity":Number(placeQuantity),"IsComputer":isComputer,"IDBuilding":Number(radioValue),"Name": number};
    requestFunc(url="/classroom/", method="POST", data=raw)
    alert("Добавлено")
}