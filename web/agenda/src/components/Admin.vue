<script setup>
import axios from "axios"
import { getTransitionRawChildren } from "vue";
</script>

<template>
	<v-form>
		<v-row>
			<v-col cols="12" md="3">
				<h4>Users</h4>
				<v-card>
					<v-list max-height="300" density="compact" :items="users" item-title="name" item-value="id" @click:select="roleUserListClick"></v-list>
				</v-card>
			</v-col>
			<v-col cols="12" md="5">
				<v-form>
					<v-row>
						<v-col cols="12" md="2">
							<v-btn @click="roleUserNewClick" color="blue">New</v-btn>
						</v-col>
						<v-col cols="12" md="2">
							<v-btn @click="roleUserSaveClick" color="teal" :disabled="disableUserRoleSave">Save</v-btn>
						</v-col>
					</v-row>
					<v-row>
						<v-col cols="12" md="6">
							<v-select label="role" variant="outlined" density="compact" :items="roles" item-title="name" item-value="id" v-model="role"></v-select>
							<v-text-field label="name" variant="outlined" density="compact" v-model="name"></v-text-field>
							<v-text-field label="password" variant="outlined" density="compact" v-model="pwd"></v-text-field>
						</v-col>
					</v-row>
				</v-form>
			</v-col>
			<v-col cols="12" md="4">
				<v-form>
					<v-row>
						<v-col>
							<h4>Defaults</h4>	
							<v-text-field label="Bishop" variant="outlined" density="compact" v-model="bishop_default" @blur="defaultUpdateBishop"></v-text-field>
							<v-text-field label="1st Counselor" variant="outlined" density="compact" v-model="bishop_1st_default" @blur="defaultUpdate1st"></v-text-field>
							<v-text-field label="2nd Counselor" variant="outlined" density="compact" v-model="bishop_2nd_default" @blur="defaultUpdate2nd"></v-text-field>
							<v-text-field label="Chorister" variant="outlined" density="compact" v-model="chorister_default" @blur="defaultUpdateChorister"></v-text-field>
							<v-text-field label="Organist" variant="outlined" density="compact" v-model="organist_default" @blur="defaultUpdateOrganist"></v-text-field>
							<v-text-field label="Stake Rep" variant="outlined" density="compact" v-model="stake_default" @blur="defaultUpdateStake"></v-text-field>
							<v-text-field label="Newsletter" variant="outlined" density="compact" v-model="newsletter_default" @blur="defaultUpdateNewsletter"></v-text-field>
							<v-text-field label="Stake Pres" variant="outlined" density="compact" v-model="stake_pres_default" @blur="defaultUpdateSPres"></v-text-field>
							<v-text-field label="Stake 1st Counselor" variant="outlined" density="compact" v-model="stake_1st_default" @blur="defaultUpdateS1st"></v-text-field>
							<v-text-field label="Stake 2nd Counselor" variant="outlined" density="compact" v-model="stake_2nd_default" @blur="defaultUpdateS2nd"></v-text-field>
						</v-col>
					</v-row>
				</v-form>
			</v-col>
		</v-row>
	</v-form>
</template>
<script>
export default {
	name: "Admin",
	data() {
		return {
			users: [],
			roles: [],
			new: false,
			name: "",
			pwd: "",
			role: "",
			id: "",
			bishop_default: "",
			bishop_1st_default: "",
			bishop_2nd_default: "",
			chorister_default: "",
			organist_default: "",
			stake_default: "",
			newsletter_default: "",
			stake_pres_default: "",
			stake_1st_default: "",
			stake_2nd_default: "",
		}
	},
	mounted() {
		this.getUserRoles();
		this.getRoles();
		this.getDefaults();
	},
	methods: {
		getUserRoles: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/roleuser/search")
            .then((response) => {
				this.users = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		getRoles: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/role/search")
            .then((response) => {
				this.roles = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		roleUserListClick: function(item) {
			this.new = false;
			axios.get(import.meta.env.VITE_API_URL + "/v1/roleuser/"+item.id)
            .then((response) => {
				this.id = item.id;
				const user = response.data.data;
				this.name = user.name;
				this.pwd = user.pwd;
				this.roles.forEach(role => {
					if (role.id === user.role_id) {
						this.role = role.name;
					}
				})
			})
			.catch(error => {
				console.log(error);
			})
		},
		roleUserNewClick: function() {
			this.new = true;
			this.name = "";
			this.pwd = "";
			this.role = "";
		},
		roleUserSaveClick: function() {
			if (this.new) {
				const obj = {name: this.name, pwd: this.pwd, role_id: this.role}
				axios.post(import.meta.env.VITE_API_URL + "/v1/roleuser", obj)
				.then(() => {
					this.getUserRoles();
					this.new = false;
					this.name = "";
					this.pwd = "";
					this.role = "";
				})
				.catch(error => {
					console.log(error);
				})
			} else {
				const obj = {id: this.id, name: this.name, pwd: this.pwd}
				axios.patch(import.meta.env.VITE_API_URL + "/v1/roleuser", obj)
				.then(() => {
					this.getUserRoles();
					this.new = false;
					this.name = "";
					this.pwd = "";
					this.role = "";
				})
				.catch(error => {
					console.log(error);
				})
			}
		},
		getDefaults: function() {
			axios.get(import.meta.env.VITE_API_URL + "/v1/defaultcalling/1")
            .then((response) => {
				const b = response.data.data;
				this.bishop_default = b.bishop;
				this.bishop_1st_default = b.b_1st;
				this.bishop_2nd_default = b.b_2nd;
				this.chorister_default = b.chorister;
				this.organist_default = b.organist;
				this.newsletter_default = b.newsletter;
				this.stake_default = b.stake;
				this.stake_pres_default = b.s_pres;
				this.stake_1st_default = b.s_1st;
				this.stake_2nd_default = b.s_2nd;
			})
			.catch(error => {
				console.log(error);
			})
		},
		defaultUpdate: function(obj) {
			obj["id"] = 1;
			axios.patch(import.meta.env.VITE_API_URL + "/v1/defaultcalling", obj)
			.catch(error => {
				console.log(error);
			})
		},
		defaultUpdateBishop: function() {
			const obj = {"bishop": this.bishop_default}
			this.defaultUpdate(obj);
		},
		defaultUpdate1st: function() {
			const obj = {"b_1st": this.bishop_1st_default}
			this.defaultUpdate(obj);
		},
		defaultUpdate2nd: function() {
			const obj = {"b_2nd": this.bishop_2nd_default}
			this.defaultUpdate(obj);
		},
		defaultUpdateChorister: function() {
			const obj = {"chorister": this.chorister_default}
			this.defaultUpdate(obj);
		},
		defaultUpdateOrganist: function() {
			const obj = {"organist": this.organist_default}
			this.defaultUpdate(obj);
		},
		defaultUpdateNewsletter: function() {
			const obj = {"newsletter": this.newsletter_default}
			this.defaultUpdate(obj);
		},
		defaultUpdateStake: function() {
			const obj = {"stake": this.stake_default}
			this.defaultUpdate(obj);
		},
		defaultUpdateSPres: function() {
			const obj = {"s_pres": this.stake_pres_default}
			this.defaultUpdate(obj);
		},
		defaultUpdateS1st: function() {
			const obj = {"s_1st": this.stake_1st_default}
			this.defaultUpdate(obj);
		},
		defaultUpdateS2nd: function() {
			const obj = {"s_2nd": this.stake_2nd_default}
			this.defaultUpdate(obj);
		}
	},
	computed: {
		disableUserRoleSave() {
			return (this.name !== "" && this.pwd !== "" && this.role !== "") ? false : true;
		}
	}
}
</script>