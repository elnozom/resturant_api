


DROP PROC IF EXISTS GroupTablesListAll
GO
CREATE  PROCEDURE GroupTablesListAll
AS
BEGIN
	SELECT GroupTableNo , GroupTableName , StartNo , (TableCount + StartNo) EndNo  FROM GroupTables
END


DROP PROC IF EXISTS GroupTablesEditAdd
GO
CREATE  PROCEDURE GroupTablesEditAdd(
    @GroupTableNo INT = 0,
    @GroupTableName VARCHAR(200) ,
    @StartNo INT ,
    @TableCount INT 
)
AS
BEGIN
	IF @GroupTableNo != 0
    BEGIN
        --this means we need to delete the group and re insert it
        DELETE FROM GroupTables WHERE GroupTableNo = @GroupTableNo
        DELETE FROM Tables WHERE GroupTableNo = @GroupTableNo
    END
	INSERT INTO GroupTables 
        (
            GroupTableNo ,
            GroupTableName ,
            StartNo ,
            TableCount 
        )  
    VALUES (
        @GroupTableNo ,
        @GroupTableName ,
        @StartNo ,
        @TableCount
    )


    SET @GroupTableNo = SCOPE_IDENTITY()


    DECLARE @loop_counter INT
    SET @loop_counter = 0
    WHILE @loop_counter < @TableCount
    BEGIN
        INSERT INTO Tables (TableNo , GroupTableNo) VALUES ((@StartNo + @loop_counter) , @GroupTableNo)
        SET @loop_counter = @loop_counter + 1
    END

    SELECT GroupTableNo insertedGroupTable
    
END
