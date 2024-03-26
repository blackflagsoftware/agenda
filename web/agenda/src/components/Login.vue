<script setup>
import { ref } from 'vue';
import axios from "axios"
</script>

<template>
	<v-form>
		<v-row>
			<v-col cols="12" md="2">
				<v-text-field id="user-input" label="User" variant="outlined" v-model="user" density="compact"></v-text-field>
			</v-col>
			<v-col cols="12" md="2">
				<v-text-field id="pwd-input" type="password" label="Password" variant="outlined" v-model="pwd" density="compact"></v-text-field>
			</v-col>
			<v-col cols="12" md="1">
				<v-btn @click="login" :disabled="disableLoginBtn" color="blue">Login</v-btn>
			</v-col>
			<v-col cols="12" md="1">
				<v-btn @click="logout" :disabled="disableLogoutBtn" color="red">Logout</v-btn>
			</v-col>
			<v-col cols="12" md="2">
				<v-alert v-if="error" type="error" text="Login Failed!"></v-alert>
			</v-col>
		</v-row>
	</v-form>
</template>

<script>
export default {
	name: "Login",
	data() {
		return {
			user: "",
			pwd: "",
			error: ref(false),
			localRole: "",
		}
	},
	props: [
		"role"
	],
	emit: {
		logout: null,
	},
	methods: {
		login: function() {
			this.error = false;
			let user = this.user.toLowerCase()
			axios.get(import.meta.env.VITE_API_URL + "/v1/roleuser/login/" + user + "/pwd/" + this.pwd)
			.then(response => {
				this.$emit("captureRole", response.data.data.role)
				this.user = "";
				this.pwd = "";
			})
			.catch(error => {
				console.log(error);
				this.error = true;
			})
		},
		logout: function() {
			this.$emit("logout");
		}
	},
	computed: {
		disableLoginBtn() {
			return (this.user === "" || this.pwd === "") ? true : false;
		},
		disableLogoutBtn() {
			return (this.localRole === "") ? true : false;
		}
	},
	watch: {
		role: {
			handler(newRole, oldRole) {
				this.localRole = newRole;
			},
			immediate: true
		}
	}
}
</script>