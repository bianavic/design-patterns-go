-- Create tables
CREATE TABLE MusicAlbum (
                            ID INT PRIMARY KEY,
                            Name VARCHAR(255),
                            MediaType VARCHAR(255),
                            Genre VARCHAR(255),
                            Details TEXT,
                            Country VARCHAR(255),
                            ReleaseDate DATE
);

CREATE TABLE Television (
                            ID INT PRIMARY KEY,
                            Name VARCHAR(255),
                            MediaType VARCHAR(255),
                            Genre VARCHAR(255),
                            Details TEXT,
                            Country VARCHAR(255),
                            ReleaseDate DATE
);

CREATE TABLE Person (
                        ID INT PRIMARY KEY,
                        Name VARCHAR(255),
                        Address TEXT,
                        Country VARCHAR(255),
                        ZipCode VARCHAR(20),
                        Email VARCHAR(255),
                        Active INT
);

CREATE TABLE Music (
                      ID INT PRIMARY KEY,
                      MusicAlbumID INT,
                      Year INT,
                      PersonID INT,
                      Title VARCHAR(255),
                      Artist VARCHAR(255),
                      FOREIGN KEY (MusicAlbumID) REFERENCES MusicAlbum(ID),
                      FOREIGN KEY (PersonID) REFERENCES Person(ID)
);

CREATE TABLE TVShow (
                        ID INT PRIMARY KEY,
                        TelevisionID INT,
                        Year INT,
                        PersonID INT,
                        Title VARCHAR(255),
                        Director VARCHAR(255),
                        FOREIGN KEY (TelevisionID) REFERENCES Television(ID),
                        FOREIGN KEY (PersonID) REFERENCES Person(ID)
);

CREATE TABLE Entertainment (
                               Media VARCHAR(255),
                               MediaType VARCHAR(255),
                               Type VARCHAR(255),
                               Description TEXT
);

-- Insert sample data
INSERT INTO MusicAlbum (ID, Name, MediaType, Genre, Details, Country, ReleaseDate) VALUES
                                                                                       (1, 'Folklore', 'CD', 'Indie folk', 'Folklore is the eighth studio album by the American singer-songwriter Taylor Swift', 'USA', '2020-07-24'),
                                                                                       (2, 'TrustFall', 'Vinyl', 'Pop', 'Trustfall is the ninth studio album by American singer Pink', 'USA', '2023-02-17');

INSERT INTO Television (ID, Name, MediaType, Genre, Details, Country, ReleaseDate) VALUES
                                                                                       (1, 'Voice', 'Streaming', 'Drama', 'Moo Jin-hyuk (Jang Hyuk) is a "mad dog" detective who becomes guilt-ridden after his wife was murdered while he was at work. Kang Kwon-joo (Lee Ha-na) is a tough policewoman who is gifted with perfect psycho-acoustics skills and went for voice profiling', 'Korea', '2019-06-11'),
                                                                                       (2, 'Beyond Evil', 'Cable', 'Comedy', 'Beyond Evil follows the story of two fearless policemen from the Manyang Police Substation[a] of the Munju Police Station (located in the western part of Gyeonggi Province), Lee Dong-sik (Shin Ha-kyun) and Han Joo-won (Yeo Jin-goo), who break the law to catch a serial killer', 'Korea', '2021-04-19');

INSERT INTO Person (ID, Name, Address, Country, ZipCode, Email, Active) VALUES
                                                                            (1, 'John Doe', '123 Main St', 'USA', '12345', 'john.doe@example.com', 1),
                                                                            (2, 'Jane Smith', '456 Elm St', 'UK', '67890', 'jane.smith@example.com', 1);

INSERT INTO Music (ID, MusicAlbumID, Year, PersonID, Title, Artist) VALUES
                                                                           (1, 1, 2020, 1, 'Exile', 'Taylor Swift'),
                                                                           (2, 2, 2023, 2, 'TrustFall', 'Pink');

INSERT INTO TVShow (ID, TelevisionID, Year, PersonID, Title, Director) VALUES
                                                                           (1, 1, 2019, 1, '16 Episodes', 'Kim Hong-Sun'),
                                                                           (2, 2, 2021, 2, '16 Episodes', 'Shim Na-yeon');

INSERT INTO Entertainment (Media, MediaType, Type, Description) VALUES
                                                                    ('Album', 'Music', 'All', 'About music'),
                                                                    ('TV Show', 'Television', 'All', 'About television');