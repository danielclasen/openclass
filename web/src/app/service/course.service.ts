import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Course} from "../model/course";

@Injectable({
  providedIn: 'root'
})
export class CourseService {

  constructor(private http: HttpClient) {}

  getCourses() {
    return this.http.get<Course[]>("/api/v1/courses")
  }

}
