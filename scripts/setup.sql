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


INSERT INTO parameters (createDate, k, m, n) VALUES ('2019-10-19 19:00:00', 3, 6 ,9);
INSERT INTO lists (createDate, IsWhite, IpAddress) VALUES ('2019-10-19 19:00:00', true, '127.0.0.1/24');