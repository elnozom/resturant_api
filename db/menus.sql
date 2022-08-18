
DROP PROC IF EXISTS MenusEditAdd
GO
CREATE  PROCEDURE MenusEditAdd (@name  VARCHAR(200) , @items TEXT)
AS
BEGIN
    DECLARE @menuSerial INT
	INSERT INTO Menus (MnuName) VALUES (@name)
	SET @menuSerial = SCOPE_IDENTITY()
	INSERT INTO ItemMenuMap (MnuSerial , ItemSerial , ItemPrice) SELECT  @menuSerial , id , price FROM  dbo.ExtractMenuItesmFromCSV(@items , ',' , '-')
	SELECT @menuSerial menuSerial
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
	@groupCode INT
)
AS
BEGIN
	declare @baseQuery VARCHAR(600) , @inMenuQuery VARCHAR(600) , @notInMenuQuery VARCHAR(600)
    SET @baseQuery = '
    SELECT i.Serial , i.ItemName ,  g.GroupCode , g.GroupName , i.BarCode , id.POSPP 
        FROM StkMs01 i
        JOIN GroupCode g ON i.GroupCode = g.GroupCode
        JOIN StkMs02 id ON i.Serial = id.ItemSerial
        LEFT JOIN ItemMenuMap im ON i.Serial = im.ItemSerial
        WHERE  g.GroupCode = dbo.ISZERO(@groupCode , g.GroupCode) AND im.Serial 
    '
    SET @inMenuQuery = CONCAT(@baseQuery , '=' , @menuId)
    SET @notInMenuQuery = CONCAT(@baseQuery , 'IS NULL')

    EXEC @inMenuQuery
    EXEC @notInMenuQuery
END
