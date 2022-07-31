USE RMSS



DROP PROC IF EXISTS GroupCodeList
GO
CREATE  PROCEDURE GroupCodeList (@parentCode  VARCHAR(8) = '')
AS
BEGIN
	
    SELECT gc.GroupCode ,  gc.GroupName , ISNULL(gc.GroupNameEn  ,gc.GroupName) GroupNameEn , gc.code
        FROM GroupCode gc
       WHERE gc.parent_code =  dbo.ISEMPTY(@parentCode , gc.parent_code)
       AND (LEN(gc.code) / 2 - 1) < 2
END


DROP PROC IF EXISTS GroupCodeFind
GO
CREATE  PROCEDURE GroupCodeFind (@groupCode  INT )
AS
BEGIN
	
    SELECT gc.GroupCode ,  gc.GroupName , ISNULL(gc.GroupNameEn  ,gc.GroupName) GroupNameEn, ISNULL(gc.parent_code , '') parent_code , gc.code
        FROM GroupCode gc
       WHERE gc.GroupCode = @groupCode
END



DROP PROC IF EXISTS GroupCodeListHierarchy
GO
CREATE  PROCEDURE GroupCodeListHierarchy
AS
BEGIN
	WITH groups (childrenLength ,  code )
	AS (
		SELECT COUNT(c.parent_code) , p.code , (LEN(p.code) / 2 - 1) groupLevel  FROM GroupCode p JOIN GroupCode c ON p.code = c.parent_code GROUP BY p.code
	)

    SELECT gc.GroupCode ,  gc.GroupName , ISNULL(gc.GroupNameEn  ,gc.GroupName) GroupNameEn, ISNULL(gc.parent_code , '') parent_code , gc.code , g.childrenLength , g.groupLevel 
        FROM GroupCode gc
        LEFT JOIN groups g ON  g.code = gc.code
    ORDER BY g.groupLevel
END



DROP PROC IF EXISTS GroupCodeInsertUpdate
GO
CREATE  PROCEDURE GroupCodeInsertUpdate (
    @groupCode  INT = 0,
    @groupName  VARCHAR(250),
    @groupNameEn  VARCHAR(250),
    @imagePath  VARCHAR(250),
    @parentCode  VARCHAR(250)
)
AS
BEGIN
    IF @groupCode = 0
        INSERT GroupCode (GroupName , GroupNameEn , parent_code ,ImagePath, code) 
        VALUES (@groupName , @groupNameEn , @parentCode , @imagePath,dbo.GROUPCODEGENERATE(@parentCode))
    ELSE
        BEGIN
            UPDATE GroupCode 
                SET GroupName = @groupName ,
                GroupNameEn = @groupNameEn ,
                ImagePath = @imagePath ,
                parentCode = @parentCode ,
                code = dbo.GROUPCODEGENERATE(@parentCode)  
            WHERE GroupCode = @groupCode
        END
    
    SELECT 
        GroupCode GroupCode ,
        GroupName GroupName ,
        ISNULL(GroupNameEn , GroupName) GroupNameEn ,
        ISNULL(parent_code , '') ,
        code 
    FROM GroupCode 
    WHERE GroupCode = dbo.ISZERO(@groupCode , SCOPE_IDENTITY())
END


