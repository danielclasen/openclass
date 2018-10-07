import {Component, OnInit} from '@angular/core';
import {CourseService} from "./service/course.service";
import {Course} from "./model/course";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

  courses: Course[];

  constructor(private courseService: CourseService) {
  }


  ngOnInit(): void {
    this.courseService.getCourses()
      .subscribe((data: Course[]) => this.courses = [...data]);

  }

}
