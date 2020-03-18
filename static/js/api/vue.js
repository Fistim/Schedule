
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
            buildings: []
         },
        methods: {
            showClassrooms() {
                console.log(this.classrooms)
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

