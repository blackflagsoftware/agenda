<script setup>
import Login from './components/Login.vue'
import Date from './components/Date.vue'
import Persons from './components/Persons.vue'
import Hymns from './components/Hymns.vue'
import Prayers from './components/Prayers.vue'
import Business from './components/Business.vue'
import Program from './components/Program.vue'
import Announcement from './components/Announcement.vue'
import Member from './components/Member.vue'
import Admin from './components/Admin.vue'
import axios from "axios"
</script>

<template>
  <v-app>
    <div style="margin-left:20px; margin-top:20px;">
      <Login @capture-role="captureRole" :role="role" @logout="logout"/>
    </div>
    <div v-if="hideShowMainDiv" style="margin-left:20px; margin-right:20px;">
        <Date @capture-agenda="captureAgenda"/>
    </div>
    <div v-if="hideShowAnnouncementDiv" style="margin-left:20px; margin-right:20px;">
        <h3 style="margin-bottom:10px; margin-top:20px;">Announcements</h3>
        <v-expansion-panels>
            <v-expansion-panel title="Announcements">
                <v-expansion-panel-text>
                    <div style="margin-top:20px">
                        <Announcement />
                    </div>
                </v-expansion-panel-text>
            </v-expansion-panel>
        </v-expansion-panels>
    </div>
    <div v-if="hideShowDetailDiv" style="margin-left:20px; margin-right:20px;">
        <h3 style="margin-bottom:10px; margin-top:20px;">Agenda/Program</h3>
        <v-expansion-panels>
          <v-expansion-panel title="Persons">
            <v-expansion-panel-text>
                <div style="margin-top:20px">
                    <Persons :agenda="agenda" @refresh-agenda="refreshAgenda"/>
                </div>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel title="Hymns">
            <v-expansion-panel-text>
                <div style="margin-top:20px">
                    <Hymns :agenda="agenda" @refresh-agenda="refreshAgenda"/>
                </div>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel title="Prayers">
            <v-expansion-panel-text>
                <div style="margin-top:20px">
                    <Prayers :agenda="agenda" @refresh-agenda="refreshAgenda"/>
                </div>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel title="Program">
            <v-expansion-panel-text>
                <div style="margin-top:20px">
                    <Program :agenda="agenda" @refresh-agenda="refreshAgenda"/>
                </div>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel title="Business">
            <v-expansion-panel-text>
                <Business :agenda="agenda" @refresh-agenda="refreshAgenda"/>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>
        <div style="margin-top:20px; margin-bottom:20px;">
            <v-btn @click="printAgenda()" style="margin-right:10px;" color="blue">Agenda Download</v-btn>
            <v-btn @click="publishProgram()" color="teal">Publish Program</v-btn>
        </div>
    </div>
    <div v-if="hideShowMainDiv" style="margin-left: 20px; margin-right:20px;">
        <h3 style="margin-bottom:10px; margin-top:20px;">Members</h3>
        <v-expansion-panels>
            <v-expansion-panel title="Members">
                <v-expansion-panel-text>
                    <Member />
                </v-expansion-panel-text>
            </v-expansion-panel>
        </v-expansion-panels>
    </div>
    <div v-if="hideShowAdmin" style="margin-left: 20px; margin-right:20px;">
        <h3 style="margin-bottom:10px; margin-top:20px;">Admin</h3>
        <v-expansion-panels>
            <v-expansion-panel title="Admin">
                <v-expansion-panel-text>
                    <Admin />
                </v-expansion-panel-text>
            </v-expansion-panel>
        </v-expansion-panels>
    </div>
    <div>
        <v-alert v-model="showAlert" :text="alertText" type="error" density="compact" closable max-width="500"></v-alert>
    </div>
  </v-app>
</template>

<script>
export default {
    name: "App",
    data() {
        return {
            role: "",
            date: "",
            agenda: undefined,
            showAlert: false,
            alertText: "error here",
        }
    },
    components: {
        Login
    },
    methods: {
        onClick: function() {
            this.showAlert = true;
        },
        logout: function() {
			sessionStorage.removeItem("role");
            this.hideShowMainDiv = false;
            this.hideShowDetailDiv = false;
            this.hideShowAnnouncementDiv = false;
            this.role = "";
        },
        captureRole: function(role) {
			sessionStorage.setItem("role", role);
            this.role = role;
        },
        captureAgenda: function(agenda) {
            this.date = agenda.date;
            this.fast_sunday = agenda.fast_sunday;
            this.agenda = agenda;
        },
        refreshAgenda: function(obj) {
            axios.patch(import.meta.env.VITE_API_URL + "/v1/agenda", obj)
            .then(() => {
                axios.get(import.meta.env.VITE_API_URL + "/v1/agenda/" + this.date)
                .then(response => {
                    this.agenda = response.data.data
                    console.log(response);
                }).catch(error => {
                    console.log(error);
                });
            })
            .catch(error => {
                console.log(error);
            })
        },
        printAgenda: function() {
            axios.get(import.meta.env.VITE_API_URL + "/v1/agenda/print/" + this.date)
            .then(() => {
                axios.get(import.meta.env.VITE_API_URL + "/documents/"+this.date+"-agenda.pdf", {responseType: "arraybuffer"})
                .then(response => {
                    let blob = new Blob([response.data], {type:'application/pdf'});
                    let link = document.createElement('a');
                    link.href = window.URL.createObjectURL(blob);
                    link.download = this.date+"-agenda.pdf";
                    link._target = 'blank';
                    link.click();
                })
                .catch(error => {
                    console.log(error);
                })
            }).catch(error => {
                console.log(error);
            });
        },
        publishProgram: function() {
            axios.get(import.meta.env.VITE_API_URL + "/v1/agenda/publish/" + this.date)
            .then(() => {
                    const date = new window.Date();
                axios.get(import.meta.env.VITE_API_URL + "/documents/"+this.date+"-program.pdf?v=" + date, {responseType: "arraybuffer"})
                .then(response => {
                    let blob = new Blob([response.data], {type:'application/pdf'});
                    let link = document.createElement('a');
                    link.href = window.URL.createObjectURL(blob);
                    link.download = this.date+"-program.pdf";
                    link._target = 'blank';
                    link.click();
                })
                .catch(error => {
                    console.log(error);
                })
            }).catch(error => {
                console.log(error);
            });
        }
    },
    computed: {
        hideShowMainDiv() {
            return this.role === "admin" || this.role === "bishopric" ? true : false;
        },
        hideShowDetailDiv() {
            return this.date === "" ? false : true;
        },
        hideShowAdmin() {
            return this.role === "admin" ? true : false;
        },
        hideShowAnnouncementDiv() {
            return this.role === "auxiliary" || this.role === "admin" || this.role === "bishopric" ? true : false;
        }
    }
}
</script>