Topics
--
- Docker compose and Configs
- CI/CD Concept and example
- Code First for Database ORM
- EOS
- Application Security Guildline


Docker
--
- Container Ecosystem/Platform
- Isolation
- Deploy on many environments
- Security
- One magic command
    ```
    https://github.com/ThanapatFox/knowledge-sharing-2026
    ```

    ``` bash
    cd db
    ```

    ``` bash
    docker compose up -d
    ```
- Check command
    - See running containers
        ``` bash
        docker ps
        ```
    - Resources monitor
        ``` bash
        docker stats
        ```


CI/CD
--
- Automation after push
- Build, Test, Pack and save package
- Workers
- Dev's responsibilities scope
    - Code Quality
    - Unit tests
    - Code Optimization


ORM
--
- Tools for Code First Practice
- Generate Schema from Code model
- Init Data from code
- Version Control


Migration Guide
--
1. Plan
    -  Known list of what to move ( Database Objects )
    - Custom Feature from old DB ( Masking Data, Any Custom View / User )
    - Type Conversion ( NVARCHAR, VARCHAR, LONG, INT, DATE format )
    - Incompatable Data Handle ( Clean Dirty Data )
2. Schema Convert
    - User/Role
    - Table and Keys
    - Skip Index for now
    - Convert SP, Trigger and Functions

3. Data Migration
    - Disable Triggers
    - Bulk Insert (tools or command for large number of rows)
    - Reset Sequences Column to match max id after imported

4. Application Connection Side
    - Update connection string
    - Switch Database engine
    - Fix Syntax ( "string_or_text_value" -> 'string_or_text_value' , SELECT [Name] FROM [TableName] -> SELECT "Name" FROM "TableName" )
    - Test for Query and Case Sensitivity

5. Verify
    - Check Objects (Tables, Views, SP, etc.)
    - Count Rows
    - Random pick complex rows to verify data
    - Re enable Keys, indexs and Triggers
    - Test


EOS
--
- Expiry Date
- Security Patches
- Risk of Using expired libs
- Check for update


Good Example
--
- Google
    - One version rule : Single version of lib across all project. Dedicated team to update version of lib version for all project.
- Netflix
    - Paved-Road : Platform team update security patches to base docker image. Application Teams use latest secure docker image for references. When Platform team push docker image will auto rebuild all Application team's containers.
- Microsoft/GitHub
    - Dependabot/Renovate : Bot to auto scan for new update in packages or requirements list and compare with CVEs database. Auto change version and do CI. 


Secure Coding
--
- OWASP Guideline
- Customer's pick guideline

References
--
- OWASP ( https://owasp.org/www-project-secure-coding-practices-quick-reference-guide/ )
- OWASP Check list ( https://owasp.org/www-project-secure-coding-practices-quick-reference-guide/stable-en/02-checklist/05-checklist.html )
- Docker compose documents ( https://docs.docker.com/reference/cli/docker/compose/ )
- Universal EOS Check ( https://endoflife.date/ )
- Code First
    - Code First EF ( https://learn.microsoft.com/en-us/ef/ef6/modeling/code-first/workflows/new-database )
    - Code First Example ( https://www.entityframeworktutorial.net/code-first/simple-code-first-example.aspx )
    - Code First Go Language using GORM Auto migration ( https://gorm.io/docs/migration.html )