USE RMSS


UPDATE GroupCode  SET ImagePath = CONCAT(GroupTypeID , '/' , GroupCode , '/' , 'Default.jpg')

UPDATE StkMs01  SET ImagePath = 
CONCAT((SELECT g.GroupTypeID
FROM GroupCode g
WHERE g.GroupCode = StkMs01.GroupCode), '/', GroupCode, '/' , BarCode , '.jpeg')
SELECT GroupTypeID
FROM GroupType gt


IF OBJECT_ID('NozAccGuests', 'U') IS NULL
BEGIN
    CREATE TABLE NozAccGuests
    (
        GuestSerial int IDENTITY(1,1) NOT NULL,
        DeviceId VARCHAR(100) NOT NULL,
        GuestName VARCHAR(100) NOT NULL,
        GuestPhone VARCHAR(100) NOT NULL,
        CreatedAt DATETIME DEFAULT GETDATE()
    )
END

IF OBJECT_ID('NozTrCart', 'U') IS NULL
BEGIN
    CREATE TABLE NozTrCart
    (
        CartSerial int IDENTITY(1,1) NOT NULL,
        TableSerial int NOT NULL,
        Amount real DEFAULT 0,
        CustomerSerial int DEFAULT NULL,
        DeviceId VARCHAR(100) NOT NULL,
        CreatedAt DATETIME DEFAULT GETDATE()
    )
END
IF OBJECT_ID('NozTrCartItems', 'U') IS NULL
BEGIN
    CREATE TABLE NozTrCartItems
    (
        CartItemSerial int IDENTITY(1,1) NOT NULL,
        CartSerial int NOT NULL,
        ItemSerial int NOT NULL,
        MainModSerial int DEFAULT 0,
        IsMod int DEFAULT 0,
        AddItems TEXT DEFAULT '',
        Qnt int NOT NULL DEFAULT 1,
        Price real NOT NULL
    );
END

GO
EXEC DropProcIfExist @Name = "CartCreate"
GO
CREATE PROC CartCreate
    (@CustomerSerial INT = 0 ,
    @DeviceId VARCHAR(100) ,
    @TableSerial INT)
AS
BEGIN
    INSERT INTO NozTrCart
        (CustomerSerial , DeviceId , TableSerial)
    VALUES
        (@CustomerSerial , @DeviceId , @TableSerial)

    SELECT SCOPE_IDENTITY() AS "Serial"
END

GO
EXEC DropProcIfExist @Name = "CartClose"
GO
CREATE PROC CartClose
    (@Serial int)
AS
BEGIN
    DECLARE @Amount real
    SELECT @Amount = SUM(Qnt * Price)
    FROM NozTrCartItems
    WHERE CartSerial = @Serial
    UPDATE NozTrCart SET Amount = @Amount WHERE CartSerial = @Serial
    SELECT 1 AS Closed
END

GO
EXEC DropProcIfExist @Name = "CartItemList"
GO
CREATE PROC CartItemList
    (@Table INT ,
    @DeviceId VARCHAR(100))
AS
BEGIN
    SELECT c.CartSerial , ci.CartItemSerial , ci.Qnt , ISNULL(Price , 0) Price ,
        ci.ItemSerial , i.ItemName , ci.IsMod, ISNULL(ci.MainModSerial , 0) MainModSerial
	  , ISNULL(ci.AddItems , '') AddItems
    FROM NozTrCart c
        JOIN NozTrCartItems ci ON c.CartSerial = ci.CartSerial
        JOIN StkMs01 i ON ci.ItemSerial = i.Serial
    WHERE c.TableSerial = @Table AND c.DeviceId = @DeviceId
    order by ci.CartItemSerial
END


GO
EXEC DropProcIfExist @Name = "CartItemCreate"
GO
CREATE PROC CartItemCreate
    (@CartSerial INT ,
    @ItemSerial INT,
    @Price REAL
)
AS
BEGIN
    INSERT INTO NozTrCartItems
        (CartSerial,ItemSerial,Price)
    VALUES
        (@CartSerial , @ItemSerial , @Price)
    SELECT SCOPE_IDENTITY() AS "Serial"
END




GO
EXEC DropProcIfExist @Name = "CartItemDelete"
GO
CREATE PROC CartItemDelete
    (@Serial INT)
AS
BEGIN
    DELETE FROM NozTrCartItems WHERE CartItemSerial = @Serial
    SELECT 1 AS Deleted

END

GO
EXEC DropProcIfExist @Name = "CartItemUpdate"
GO
CREATE PROC CartItemUpdate
    (@Serial INT ,
    @Qnt INT)
AS
BEGIN
    UPDATE NozTrCartItems SET Qnt = @Qnt WHERE CartItemSerial = @Serial
    SELECT 1 AS Updated
END


IF OBJECT_ID('NozCartCalls', 'U') IS NULL
BEGIN
    CREATE TABLE NozCartCalls
    (
        NozCartCallsSerial int IDENTITY(1,1) NOT NULL,
        CartSerial int NOT NULL,
        TableSerial int NOT NULL,
        CreatedAt DATETIME DEFAULT GETDATE(),
        RespondedAt DATETIME,
        WaiterCode INT,
        CallType SMALLINT
    )
END


GO
EXEC DropProcIfExist @Name = "CartCallCreate"
GO
CREATE PROC CartCallCreate
    (@CallType BIT ,
    @CartSerial INT ,
    @TableSerial INT ,
    @GuestName VARCHAR(100)
    )
AS
BEGIN
    IF EXISTS (SELECT NozCartCallsSerial FROM NozCartCalls WHERE CartSerial = @CartSerial AND CallType = @CallType  AND GuestName = @GuestName AND RespondedAt IS NULL )
    BEGIN
        SELECT 0 AS Created
        RETURN
    END
    IF @CallType = 1 AND NOT EXISTS (
        SELECT * FROM StkTr03 o 
            WHERE  TableSerial = @TableSerial AND 
                    ISNULL(TotalCash,0) = 0 
        )
        BEGIN
            SELECT 0 AS Created
            RETURN
    END
    IF @CallType = 1 AND  EXISTS (
        SELECT * FROM StkTr03 o 
            JOIN  StkTr04 d 
            ON o.Serial = d.HeadSerial
            WHERE  TableSerial = @TableSerial AND 
                    ISNULL(TotalCash,0) = 0  AND d.Printed = 0 
        )
        BEGIN
            SELECT 0 AS Created
            RETURN
        END

    
    INSERT INTO NozCartCalls
        (CallType , CartSerial , TableSerial , GuestName)
    VALUES
        (@CallType , @CartSerial , @TableSerial , @GuestName)
    SELECT 1 AS Created
END




GO
EXEC DropProcIfExist @Name = "CartCheckCalls"
GO
CREATE PROC CartCheckCalls
    (@Imei VARCHAR(100))
AS
BEGIN
    SELECT c.CartSerial, c.TableSerial, c.CallType ,t.GroupTableNo , t.TableNo , ISNULL(c.GuestName , '') , c.CreatedAt  
    FROM NozCartCalls c 
        JOIN
            Tables t ON c.TableSerial = t.Serial 
        JOIN
            GTE_Map g ON t.GroupTableNo = g.GTID 
        JOIN
            ComUse com ON g.EmpID = com.UserId
    WHERE com.Imei = @Imei AND RespondedAt IS NULL
END



GO
EXEC DropProcIfExist @Name = "GuestsCreate"
GO
CREATE PROC GuestsCreate
    (
    @DeviceId VARCHAR(100),
    @GuestName VARCHAR(100),
    @GuestPhone VARCHAR(100)
)
AS
BEGIN
    INSERT INTO NozAccGuests
        (
        DeviceId,
        GuestName,
        GuestPhone
        )
    VALUES
        (
            @DeviceId,
            @GuestName,
            @GuestPhone
    )

    SELECT SCOPE_IDENTITY() AS "Serial"

END



GO
EXEC DropProcIfExist @Name = "GuestsGetByDeivce"
GO
CREATE PROC GuestsGetByDeivce
    (
    @DeviceId VARCHAR(100)
)
AS
BEGIN
    SELECT GuestName, GuestPhone
    FROM NozAccGuests
    WHERE DeviceId = @DeviceId
END

GO
ALTER PROCEDURE [dbo].[StkMs01ListByMenuAndGroup](@GroupCode int,
        @TableSerial int =0)
AS
BEGIN
    DECLARE @menuSerial int
    SELECT @menuSerial = Serial
    FROM Menus
    WHERE IsTS = 1
    ;
    with
        Items (ItemSerial, ItemDesc,ImagePath, ItemPrice, ItemCode, ItemName, WithModifier)
        as
        (
            SELECT im.ItemSerial , i.ItemDesc ,i.ImagePath, im.ItemPrice , i.ItemCode , i.ItemName , i.WithModifier
            FROM ItemMenuMap  im JOIN StkMs01 i ON im.ItemSerial = i.Serial AND i.GroupCode = @GroupCode
            WHERE im.MnuSerial = @MenuSerial
        )
 ,
        SalesItems (TQnt, ItemSerial)
        as
        (
            Select sum(Qnt) , ItemSerial
            from StkTr04 inner join StkTr03 on StkTr04.HeadSerial = StkTr03.Serial
            where StkTr03.TableSerial = @TableSerial and isnull(StkTr03.TotalCash ,0) = 0
            group by StkTr04.ItemSerial
        )

    Select Items.ItemSerial, ItemPrice,ImagePath, ItemCode, ItemName, ItemDesc, WithModifier, isnull(SalesItems.TQnt,0)
    from Items
        left join SalesItems on Items.ItemSerial = SalesItems.ItemSerial


END

GO
ALTER PROCEDURE [dbo].[GroupCodeListByGroupTypeId](@GroupTypeID int)
AS
BEGIN

	SET NOCOUNT ON;
	SELECT  g.GroupCode , g.GroupName , g.ImagePath FROM  "GroupCode" g WHERE g.GroupTypeID = @GroupTypeID AND g.ShowOnMenu = 1
END

ALTER TABLE StkMs01
ADD ItemNameEn VARCHAR(100); 


ALTER TABLE GroupCode
ADD GroupImage VARCHAR(100); 


ALTER TABLE GroupCode
ADD GroupNameEn VARCHAR(100); 


ALTER TABLE GroupType
ADD GroupTypeNameEn VARCHAR(100);



ALTER TABLE StkMs01
ADD ImagePath VARCHAR(100); 


ALTER TABLE NozCartCalls
ADD GuestName VARCHAR(100); 

