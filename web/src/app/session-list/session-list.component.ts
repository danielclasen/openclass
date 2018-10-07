import {Component, Input, OnInit} from '@angular/core';
import {SessionService} from "../service/session.service";
import {Course} from "../model/course";
import {Session} from "../model/session";

@Component({
  selector: 'app-session-list',
  templateUrl: './session-list.component.html',
  styleUrls: ['./session-list.component.scss']
})
export class SessionListComponent implements OnInit {

  @Input()
  public course: Course;

  private sessions: Session[] = [];

  constructor(private sessionService: SessionService) { }

  ngOnInit() {
    this.sessionService.getSessionsForCourse(this.course.id).subscribe((data: Session[]) => this.sessions = [...data]);
  }

}
