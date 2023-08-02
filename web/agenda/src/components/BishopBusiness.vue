<script setup>
import BishopUpdate from './BishopBusinessUpdate.vue'
import axios from "axios"
</script>

<template>
	<div style="margin-bottom:10px">
		<b>Bishop Business</b>
	</div>
	<v-form>
		<v-row>
			<v-col cols="12" md="11">
				<v-text-field label="Message" variant="outlined" density="compact" v-model="addMessage"></v-text-field>
			</v-col>
			<v-col cols="12" md="1">
				<v-btn @click="addSave" :disabled="disableSaveBtn" color="blue">Add</v-btn>
			</v-col>
		</v-row>
	</v-form>
	<v-sheet v-if="showNoBishop">No Bishop Business</v-sheet>
	<v-form>
		<BishopUpdate v-for="v in bishop" :item="v" @refresh-bishop="getBishop"/>
	</v-form>
</template>

<script>
export default {
	name: "BishopBusiness",
	data() {
		return {
			addMessage: "",
			bishop: [],
		}
	},
	props: [
		"date"
	],
	mounted() {
		this.getBishop();
	},
	methods: {
		getBishop: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/bishopbusiness/search", {"search": [{"column": "date", "value": this.date, "compare": "="}]})
			.then(response => {
				this.bishop = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		addSave: function() {
			const obj = {"date": this.date, "message": this.addMessage}
			axios.post(import.meta.env.VITE_API_URL + "/v1/bishopbusiness", obj)
			.then(() => {
				this.addMessage = "";
				this.getBishop();
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableSaveBtn() {
			return (this.addMessage === "") ? true : false;
		},
		showNoBishop() {
			return this.bishop.length === 0 ? true : false;
		}
	}
}
</script>