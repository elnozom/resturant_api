
DROP PROC IF EXISTS MenusAttach
GO
CREATE  PROCEDURE MenusAttach (@id INT , @items TEXT)
AS
BEGIN
	INSERT INTO ItemMenuMap (MnuSerial , ItemSerial , ItemPrice) SELECT  @id , id , price FROM  dbo.ExtractMenuItesmFromCSV(@items , ',' , '-')
	SELECT @id id
END

DROP PROC IF EXISTS MenusList
GO
CREATE  PROCEDURE MenusList 
AS
BEGIN
    SELECT   Serial , Name , CreatedAt FROM Menus
END


DROP PROC IF EXISTS MenusEditAdd
GO
CREATE  PROCEDURE MenusEditAdd (@id INT , @name  VARCHAR(200))
AS
BEGIN
	IF @id != 0
        BEGIN
            UPDATE Menus SET MnuName = @name WHERE Serial = @id
            SELECT @id id
            RETURN
        END

	INSERT INTO Menus (MnuName) VALUES (@name)
    SELECT SCOPE_IDENTITY() id
END

GO
CREATE  PROCEDURE LoadMenuItems(@menuId INT = 0)
AS
BEGIN
	WITH groups (groupCode , groupName )
	AS (
		SELECT p.GroupCode , p.GroupName   FROM GroupCode p  WHERE (SELECT COUNT(*) FROM GroupCode WHERE parent_code = p.code) = 0
	)

	SELECT i.Serial , i.ItemName , i.GroupCode FROM StkMs01 i JOIN StkMs02 d ON i.Serial = d.ItemSerial JOIN groups g ON i.GroupCode = g.code
END

GO
ALTER FUNCTION [dbo].[ExtractMenuItesmFromCSV] (@list NVARCHAR(MAX) ,@delimiter CHAR(1) , @delimiter2 CHAR(1))
RETURNS @table TABLE ( 
     id INT , price REAL )
AS 
    BEGIN
        DECLARE @pos INT ,@nextpos INT ,@valuelen INT ,@pricePos INT, @idLen INT  ,@currentItem VARCHAR(100) , @currentId INT , @currentPrice REAL
        SELECT  @pos = 0 ,@nextpos = 1
		
		WHILE @nextpos > 0 
            BEGIN
                SET  @nextpos = CHARINDEX(@delimiter,@list,@pos + 1)
                SET  @valuelen = CASE WHEN @nextpos > 0 THEN @nextpos
                                         ELSE LEN(@list) + 1
                                    END - @pos - 1

				SET @currentItem = SUBSTRING(@list,@pos + 1,@valuelen)

				SET  @pricePos = CHARINDEX(@delimiter2,@currentItem)
				SET @idLen = @pricePos - 1
				SET @currentId = CONVERT(INT ,  SUBSTRING(@currentItem,1,@idLen))
				SET @currentPrice = CONVERT(REAL , SUBSTRING(@currentItem,@idLen + 2,10))
                INSERT  @table ( id , price )
                VALUES  (@currentId , @currentPrice) 
                SELECT  @pos = @nextpos

            END

        DELETE  FROM @table
        WHERE   id = ''

        RETURN 
    END

GO
CREATE  PROC MenuItemsList(
	@menuId INT = 0,
	@groupCode INT = 0
)
AS
BEGIN
		 SELECT im.Serial ,
                i.Serial ,
                i.ItemName ,
                g.GroupCode ,
                g.GroupName ,
                i.BarCode , 
                CASE WHEN im.Serial IS NULL 
                    THEN  id.POSPP 
                    ELSE im.ItemPrice 
                END price
            FROM StkMs01 i
            JOIN GroupCode g ON i.GroupCode = g.GroupCode
            JOIN StkMs02 id ON i.Serial = id.ItemSerial
            LEFT JOIN ItemMenuMap im 
                ON i.Serial = im.ItemSerial
                AND im.MnuSerial =  dbo.ISZERO(@menuId , im.MnuSerial)
            WHERE  g.GroupCode = dbo.ISZERO(@groupCode , g.GroupCode)
	
END




GO
CREATE  PROC MenusPriceEdit(
	@id INT,
	@price REAL
)
AS
BEGIN
		UPDATE ItemMenuMap SET ItemPrice = @price WHERE Serial = @id
        SELECT @id id
END



GO
CREATE  PROC MenusDelete(
	@id INT
)
AS
BEGIN
    UPDATE Menus SET DeletedAt = GETDATE() WHERE Serial = @id
    SELECT 1 deleted
END


