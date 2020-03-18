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
            subjects:[], 
         },
        methods: {
            showSubjects() {
                console.log(this.subjects)
            },
        	async getSubjects(){
        		this.subjects = await requestFunc("/subject/", "GET")
        	},
            
        }, 
        mounted() {
        	this.getSubjects()
        }
    });

addSubject = () =>{
    let name = document.querySelector('.subjectName').value
    let shortName = document.querySelector('.shortName').value
    var select = document.getElementById("selectProfmodule");
    var value = select.value;
    let type = document.querySelector('.type').checked

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var raw = {"Name":name,"ID_professionalmodule":Number(value),"Shortname": shortName, "IDStype": Number(type)};
    requestFunc(url="/subject/", method="POST", data=raw)
    alert("Добавлено")
}