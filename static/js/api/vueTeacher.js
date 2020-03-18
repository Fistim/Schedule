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
            selected:null,
            teachers:[],
            classrooms:[]
         },
        methods: {
            showTeachers() {
                console.log(this.teachers)
            },
        	async getTeachers(){
        		this.teachers = await requestFunc("/teacher", "GET")
        	},
            
            async getClassrooms(){
                classrooms = await requestFunc("/classroom", "GET")
                for(let i =0; i < classrooms.length; i++) {
                    Vue.set(this.classrooms, i, classrooms[i]);
                }
            },
            async mountFunc(){
                await this.getTeachers()
                await this.getClassrooms()
            },
        },

        async mounted() {
        	await this.mountFunc()
            console.log(this.classrooms)
        }
    });

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

