<script setup>
import Visitor from "./Visitor.vue";
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="3">
			<v-text-field id="presiding-input" label="Presiding" variant="outlined" v-model="presiding" density="compact"></v-text-field>
		</v-col>
		<v-col cols="12" md="3">
			<v-text-field id="conducting-input" label="Conducting" variant="outlined" v-model="conducting" density="compact"></v-text-field>
		</v-col>
		<v-col cols="12" md="3">
			<v-text-field id="stake-input" label="Stake Rep" variant="outlined" v-model="stake" density="compact"></v-text-field>
		</v-col>
	</v-row>
	<v-row>
		<v-col cols="12" md="3">
			<v-text-field id="organist-input" label="Organist" variant="outlined" v-model="organist" density="compact"></v-text-field>
		</v-col>
		<v-col cols="12" md="3">
			<v-text-field id="chorister-input" label="Chorister" variant="outlined" v-model="chorister" density="compact"></v-text-field>
		</v-col>
		<v-col cols="12" md="3">
			<v-text-field id="newsletter-input" label="Newsletter" variant="outlined" v-model="newsletter" density="compact"></v-text-field>
		</v-col>
	</v-row>
	<v-row>
		<v-col>
			<v-btn @click="save" color="blue">Save</v-btn>
		</v-col>
	</v-row>
	<v-row>
		<v-col>
			<Visitor :date="agenda.date"/>
		</v-col>
	</v-row>
</template>

<script>
export default {
	name: "Persons",
	data() {
		return {
			presiding: "",
			conducting: "",
			organist: "",
			chorister: "",
			newsletter: "",
			stake: ""
		}
	},
	props: ["agenda"],
	emits: {
		refreshAgenda: null,
	},
	methods: {
		getDefaultCallings: function() {
			axios.get(import.meta.env.VITE_API_URL + "/v1/defaultcalling/1")
			.then(response => {
				if (this.organist === null) { this.organist = response.data.data.organist; }
				if (this.chorister === null) { this.chorister = response.data.data.chorister; }
				if (this.newsletter === null) { this.newsletter = response.data.data.newsletter; }
				if (this.stake === null) { this.stake = response.data.data.stake; }
			})
			.catch(error => {
				console.log(error);
			})
		},
		save: function() {
			const obj = {"date": this.agenda.date, "presiding": this.presiding, "conducting": this.conducting, "organist": this.organist, "chorister": this.chorister, "stake": this.stake, "newsletter": this.newsletter};
			this.$emit("refreshAgenda", obj);
		}
	},
	watch: {
		agenda: {
		 	handler(newAgenda, oldAgenda) {
				this.presiding = newAgenda.presiding;
				this.conducting = newAgenda.conducting;
				this.organist = newAgenda.organist;
				this.chorister = newAgenda.chorister;
				this.newsletter = newAgenda.newsletter;
				this.stake = newAgenda.stake;
				this.getDefaultCallings();
			},
			immediate: true
		}
	}
}
</script>