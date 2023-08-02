<script setup>
import VisitorUpdate from './VisitorUpdate.vue'
import axios from "axios"
</script>

<template>
	<div style="margin-bottom:10px">
		<b>Visitors</b>
	</div>
	<v-form>
		<v-row>
			<v-col cols="12" md="11">
				<v-text-field label="Name" variant="outlined" density="compact" v-model="addName"></v-text-field>
			</v-col>
			<v-col cols="12" md="1">
				<v-btn @click="addSave" :disabled="disableSaveBtn" color="blue">Add</v-btn>
			</v-col>
		</v-row>
	</v-form>
	<v-sheet v-if="showNoVisitor">No Visitors</v-sheet>
	<v-form>
		<VisitorUpdate v-for="v in visitor" :item="v" @refresh-visitor="getVisitor"/>
	</v-form>
</template>

<script>
export default {
	name: "Visitor",
	data() {
		return {
			addName: "",
			localDate: "",
			visitor: [],
		}
	},
	props: [
		"date"
	],
	mounted() {
		this.getVisitor();
	},
	methods: {
		getVisitor: function() {
			this.visitor = [];
			axios.post(import.meta.env.VITE_API_URL + "/v1/visitor/search", {"search": [{"column": "date", "value": this.localDate, "compare": "="}]})
			.then(response => {
				this.visitor = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		addSave: function() {
			const obj = {"date": this.localDate, "name": this.addName}
			axios.post(import.meta.env.VITE_API_URL + "/v1/visitor", obj)
			.then(() => {
				this.addName = "";
				this.getVisitor();
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableSaveBtn() {
			return (this.addName === "") ? true : false;
		},
		showNoVisitor() {
			return this.visitor.length === 0 ? true : false;
		}
	},
	watch: {
		date: {
			handler(newDate, oldDate) {
				this.localDate = newDate;
				this.getVisitor();
			},
			immediate: true
		}
	}
}
</script>