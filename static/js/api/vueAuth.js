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
	delimiters: ['${', '}'],
    el: '#app',
    data: {
    	userdata: {
    		Login: null,
    		Password: null,
    	}
    },
    methods: {
    	async authButton() {
    		let res = await requestFunc("/auth/", "POST", this.userdata);
    		localStorage.setItem("Login", this.userdata.Login);
    		localStorage.setItem("Password", this.userdata.Password);
    		// console.log(Array.isArray(res))
    		this.nulldata();
    		if (Array.isArray(res)) {
    			location.href = 'http://schedule.tomtit.tomsk.ru/plan/'
    		}
    		// return Array.isArray(res)
    	},
    	nulldata() {
            this.userdata.Login = null;
            this.userdata.Password = null;
        }
    },

});
// authButton = () =>{
// 	let login = document.querySelector('.login').value
// 	let password = document.querySelector('.password').value

// }
