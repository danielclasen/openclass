import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from "@angular/common/http";
import {Person} from "../model/person";
import {Session} from "../model/session";

@Injectable({
  providedIn: 'root'
})
export class SessionService {

  constructor(private http: HttpClient) {
  }

  getSessionsForCourse(id: number){
    return this.http.get<Session[]>(`/api/v1/courses/${id}/sessions`,{params: new HttpParams().set('id', String(id))})
  }

  getParticipationsForSession(id: number) {
    return this.http.get<Person[]>(`/api/v1/sessions/${id}/participations`, {params: new HttpParams().set('id', String(id))})
  }

  participateInSession(id: number, person: Person) {
    return this.http.post(`/api/v1/sessions/${id}/participations`, person, {params: new HttpParams().set('id', String(id))})
  }

}
