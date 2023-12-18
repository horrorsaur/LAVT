export namespace valorant {
	
	export class SessionResponse {
	    name: string;
	    pid: string;
	    puuid: string;
	    state: string;
	
	    static createFrom(source: any = {}) {
	        return new SessionResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.pid = source["pid"];
	        this.puuid = source["puuid"];
	        this.state = source["state"];
	    }
	}

}

