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
var vueApp = new Vue({
        el: '#app',
        delimiters: ['${', '}'],
        data: {
            specialties:[], 
            selectedRow: {
                "Code": "",
                "Name": "",
                "IDDuration": 0,
                "ID": ""
            },
            requestType: null
         },
        methods: {
        	async getSpecialties(){
        		this.specialties = await requestFunc("/specialty", "GET")
        	},
            selectRow(row, index){
                let rows = document.querySelectorAll('.clickHover')
                this.selectedRow = row;
                // empty string
                if(rows[index].style.backgroundColor != 'red'){
                    rows[index].style.backgroundColor = 'red'
                }
                else{
                    rows[index].style.backgroundColor = 'white'
                    this.clearSelectedRow()
                }   
            },
            clearSelectedRow(){
                this.selectedRow =   {
                "Code": "",
                "Name": "",
                "IDDuration": 0
                }
            },
            changeRequestType(requestType){
                this.requestType = requestType 
            },
            modalClick(add, requestType) {
                if(add==='yes') {
                    this.clearSelectedRow();
                }
                this.changeRequestType(requestType);
            }         
        }, 
        mounted() {
        	this.getSpecialties()
        }
    });
    addSpecialty = (requestType)=>{
        let code = document.querySelector('.code').value
        let name = document.querySelector('.name').value
        let duration = document.querySelectorAll('.duration-radio');
        let radioValue
        for(let i=0; i<duration.length; i++ ){
            if (duration[i].checked) {
                radioValue = duration[i].value

            }
        }
        var myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");

        var raw = {"Code":code,"Name":name,"IDDuration":Number(radioValue)};
        if (vueApp.requestType=== 'PATCH') {
            let ID = document.querySelector('.ID').value
            console.log(ID)
            raw.ID= Number(ID)
            
        }
        requestFunc(url="/specialty/", method=vueApp.requestType, data=raw)
        alert("Добавлено")
    }