Feature: sessions api

  Scenario: List all sessions of a course
    When I send a GET request to /api/v1/courses/1/sessions
    Then the response should be 200 and match this json:
      """
      [
        {
          "id": 1,
          "courseId": 1,
          "date": "2018-10-15T14:00:00Z",
          "location": "Bonn",
          "maxStudents": 10
        },
        {
          "id": 2,
          "courseId": 1,
          "date": "2018-10-16T14:00:00Z",
          "location": "Köln",
          "maxStudents": 10
        }
      ]
      """

  Scenario: List a single session by id
    When I send a GET request to /api/v1/sessions/1
    Then the response should be 200 and match this json:
      """
      {
        "id": 1,
        "courseId": 1,
        "date": "2018-10-15T14:00:00Z",
        "location": "Bonn",
        "maxStudents": 10
      }
      """

  Scenario: Create a new session
    When I send a POST request to /api/v1/sessions with body:
      """
      {
        "id": 3,
        "courseId": 1,
        "date": "2018-10-17T14:00:00Z",
        "location": "Frechen",
        "maxStudents": 10
      }
      """
    Then the response should be 201 and match this json:
      """
      {
        "id": 3,
        "courseId": 1,
        "date": "2018-10-17T14:00:00Z",
        "location": "Frechen",
        "maxStudents": 10
      }
      """
    When I send a GET request to /api/v1/courses/1/sessions
    Then the response should be 200 and match this json:
      """
      [
        {
          "id": 1,
          "courseId": 1,
          "date": "2018-10-15T14:00:00Z",
          "location": "Bonn",
          "maxStudents": 10
        },
        {
          "id": 2,
          "courseId": 1,
          "date": "2018-10-16T14:00:00Z",
          "location": "Köln",
          "maxStudents": 10
        },
        {
          "id": 3,
          "courseId": 1,
          "date": "2018-10-17T14:00:00Z",
          "location": "Frechen",
          "maxStudents": 10
        }
      ]
      """

  @ignore
  Scenario: Update a session
    When I send a PUT request to /api/v1/sessions/1 with body:
      """
      {
        "location": "Siegburg"
      }
      """
    Then the response should be 200 and match this json:
      """
      {
        "id": 1,
        "courseId": 1,
        "date": "2018-10-15T14:00:00Z",
        "location": "Siegburg",
        "maxStudents": 10
      }
      """

  Scenario: Participate in a session
    When I send a POST request to /api/v1/sessions/1/participations with body:
      """
      {
        "firstName": "Daniel",
        "lastName": "Clasen",
        "email": "foo@bar.com"
      }
      """
    Then the response should be 204
    When I send a GET request to /api/v1/sessions/1/participations
    Then the response should be 200 and match this json:
      """
      [
        {
          "firstName": "Daniel",
          "lastName": "Clasen",
          "email": "foo@bar.com"
        }
      ]
      """