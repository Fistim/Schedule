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
            schedule: [],
            scheduleForDay: null
         },
        methods: {
            getScheduleForDay(day){
            	for(let i = 0; i < this.schedule.length;i++) {
            		scheduleDay = this.schedule[i]
            		if(scheduleDay[0].Day === day) {
            			this.scheduleForDay = scheduleDay
            			console.log(scheduleDay)

            		}
            	}
            },
            async generateSchedule() {
            	this.schedule = await requestFunc('/schedule/generate', 'GET')
            	while(true) {
            		if(this.schedule !== null) {
            			this.getScheduleForDay(1)
            			break
            		}

            	}
            }
            
        }
    });

