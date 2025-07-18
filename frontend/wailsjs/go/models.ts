export namespace internal {
	
	export class Config {
	    display_id: number;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.display_id = source["display_id"];
	    }
	}
	export class Params {
	    tm: number;
	    x: number;
	    y: number;
	    w: number;
	    h: number;
	    name: string;
	    button: string;
	    double: boolean;
	    str: string;
	    speed: number;
	    pid: number;
	    direction: string;
	    path: string;
	    color: string;
	    keys: string[];
	
	    static createFrom(source: any = {}) {
	        return new Params(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tm = source["tm"];
	        this.x = source["x"];
	        this.y = source["y"];
	        this.w = source["w"];
	        this.h = source["h"];
	        this.name = source["name"];
	        this.button = source["button"];
	        this.double = source["double"];
	        this.str = source["str"];
	        this.speed = source["speed"];
	        this.pid = source["pid"];
	        this.direction = source["direction"];
	        this.path = source["path"];
	        this.color = source["color"];
	        this.keys = source["keys"];
	    }
	}
	export class Result {
	    data: any;
	    message: string;
	    status: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.message = source["message"];
	        this.status = source["status"];
	    }
	}

}

