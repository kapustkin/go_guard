CREATE TABLE parameters (
        id SERIAL PRIMARY KEY,
        createDate timestamp,
        k INT,
        m INT,
        n INT
);

CREATE TABLE lists (
        id SERIAL PRIMARY KEY,
        createDate timestamp,
    	IsWhite BOOLEAN DEFAULT false,
	    IpAddress text
);