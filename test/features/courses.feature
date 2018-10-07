Feature: courses api

  Scenario: List courses
    When I send a GET request to /api/v1/courses
    Then the response should be 200 and match this json:
      """
      [
        {
          "id": 1,
          "title": "Java for beginner",
          "description": "A Java beginner guide.",
          "lecturer": "Mr. Foo",
          "price": 9000
        },
        {
          "id": 2,
          "title": "Golang for experts",
          "description": "A deep dive into the world of go.",
          "lecturer": "Mr. Bar",
          "price": 18000
        }
      ]
      """

  @ignore
  Scenario: List courses with sessions in a specified time range
    """
    from: 01.10.2018 00:00:00 =timestamp in millis=> 1538352000000
    till: 31.10.2018 23:59:59 =timestamp in millis=> 1541030399000
    """
    When I send a GET request to /api/v1/courses?from=1538352000000&till=1540944000000
    Then the response should be 200 and match this json:
      """
      [
        {
          "id": 1,
          "title": "Java for beginner",
          "description": "A Java beginner guide.",
          "lecturer": "Mr. Foo",
          "price": 9000
        },
        {
          "id": 2,
          "title": "Golang for experts",
          "description": "A deep dive into the world of go.",
          "lecturer": "Mr. Bar",
          "price": 18000
        }
      ]
      """


  Scenario: Create a new course
    When I send a POST request to /api/v1/courses with body:
      """
      {
          "id": 3,
          "title": "Java for experts",
          "description": "A deep dive in Java.",
          "lecturer": "Mr. Foo",
          "price": 18000
      }
      """
    Then the response should be 201 and match this json:
      """
      {
          "id": 3,
          "title": "Java for experts",
          "description": "A deep dive in Java.",
          "lecturer": "Mr. Foo",
          "price": 18000
      }
      """
    When I send a GET request to /api/v1/courses
    Then the response should be 200 and match this json:
      """
      [
        {
          "id": 1,
          "title": "Java for beginner",
          "description": "A Java beginner guide.",
          "lecturer": "Mr. Foo",
          "price": 9000
        },
        {
          "id": 2,
          "title": "Golang for experts",
          "description": "A deep dive into the world of go.",
          "lecturer": "Mr. Bar",
          "price": 18000
        },
        {
          "id": 3,
          "title": "Java for experts",
          "description": "A deep dive in Java.",
          "lecturer": "Mr. Foo",
          "price": 18000
        }
      ]
      """

  @ignore
  Scenario: Update a course
    When I send a PUT request to /api/v1/courses/1 with body:
      """
      {
          "title": "Java for experts II",
          "lecturer": "Mr. Foo & Mr. Bar"
      }
      """
    Then the response should be 200 and match this json:
      """
      {
          "id": 1,
          "title": "Java for experts II",
          "description": "A deep dive in Java.",
          "lecturer": "Mr. Foo & Mr. Bar",
          "price": 18000
      }
      """