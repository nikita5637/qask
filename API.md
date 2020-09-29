## **Get question**
----------

_Returns json data with question_
* **URL**

    >   [`GET`] /questions

* **Request**

    * **HTTP Headers**
        >    `Content-Type`: application/json  
    
    * **Request Body**

        >    `from`=[string, required]  

* **Success Response:**

    * **Code:** 200 OK  
      **Content:** `{ question: [string] }`

* **Error Response:**

  * **Code:** 400 BAD REQUEST  
    **Content:** `{ code: [integer], message: [string] }`

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ code: [integer], message: [string] }`

* **Sample Call:**

    ```
    curl -v -X GET --data '{"from": "telegram"}' 172.20.0.3:30001/questions
    ```

<br/>

## **Create new user**
----------

_Creates new user and returns json data with question_
* **URL**

    >   [`POST`] /users

* **Request**

    * **HTTP Headers**
        >    `Content-Type`: application/json  
    
    * **Request Body**

        >    `from`=[string, required]  
            `firstName`=[string, required]  
            `userName`=[string, required]  
            `tgID`=[unsigned integer 64, optional]

* **Success Response:**

    * **Code:** 201 CREATED  
    **Content:** `{}`

* **Error Response:**

     * **Code:** 400 BAD REQUEST <br />
    **Content:** `{ code: [integer], message: [string] }`

    * **Code:** 422 UNPROCESSABLE ENTRY <br />
    **Content:** `{ code: [integer], message: [string] }`

* **Sample Call:**

    ```
    curl -v -X POST --data '{"from": "telegram", "firstName": "nikita", "userName": "nikita_arbavian", "tgID": 1}' 172.20.0.3:30001/users
    ```