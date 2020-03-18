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

var vueapp = new Vue({
        el: '#app',
        delimiters: ['${', '}'],
        data: {
            groups:[],
            teachers:[],
            specialties:[],
        },
        methods: {
            showGroups() {
                console.log(this.groups)
            },
        	async getGroups(){
        		this.groups = await requestFunc("/group/", "GET")
        	},
            async getTeachers(){
                teachers = await requestFunc("/teacher/", "GET")
                for(let i =0; i < teachers.length; i++) {
                    Vue.set(this.teachers, i, teachers[i]);
                }
            },
            async getSpecilties(){
                specialties = await requestFunc("/specialty/", "GET")
                for(let i =0; i < specialties.length; i++) {
                    Vue.set(this.specialties, i, specialties[i]);
                }
            },
            async mountFunc(){
                await this.getTeachers()
                await this.getGroups()
                await this.getSpecilties()
            }
        }, 
        async mounted() {
            await this.mountFunc()
            console.log(this.teachers)
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

