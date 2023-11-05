<script setup>
import axios from "axios"
import SpeakerType from "./SpeakerType.vue";
import Hymn from "./Hymn.vue";
</script>

<template>
	<v-row>
		<v-col cols="12" md="1">
			<v-select label="Position" variant="outlined" density="compact" v-model="position" :items="['1', '2', '3', '4', '5', '6', '7', '8', '9', '10']"></v-select>
		</v-col>
		<v-col cols="12" md="2">
			<SpeakerType :value-in="speaker_type" @speaker-type-out="speakerTypeOut" /> 
		</v-col>
		<v-col cols="12" md="7">
			<Hymn :hymns="hymns" :hymn-number-in="this.item.name" @hymn-number-out="hymnNumberOut" v-if="speakerTypeEqHymn" />
			<v-text-field label="Name" variant="outlined" density="compact" v-model="name" v-else></v-text-field>
		</v-col>
		<v-col cols="12" md="1">
			<v-btn @click="updateSave()" :disabled="disableUpdateBtn" color="blue">Update</v-btn>
		</v-col>
		<v-col cols="12" md="1">
			<v-btn @click="updateDelete()" color="red">Delete</v-btn>
		</v-col>
	</v-row>
</template>

<script>
export default {
	name: "ProgramUpdate",
	data() {
		return {
			position: this.item.position,
			name: this.item.name,
			speaker_type: this.item.speaker_type,
		}
	},
	props: [
		"item",
		"hymns"
	],
	methods: {
		updateSave: function() {
			const obj = {"id": this.item.id, "position": this.position, "name": this.name}
			var url = import.meta.env.VITE_API_URL + "/v1/speaker";
			axios.patch(url, obj)
			.then(() => {
				this.$emit("refreshSpeakers");
			})
			.catch(error => {
				console.log(error);
			})
		},
		updateDelete: function() {
			var url = import.meta.env.VITE_API_URL + "/v1/speaker/"+this.item.id
			axios.delete(url)
			.then(() => {
				this.$emit("refreshSpeakers");
			})
			.catch(error => {
				console.log(error);
			})
		},
		speakerTypeOut: function(speakerType) {
			this.speaker_type = speakerType
		},
		hymnNumberOut: function(hymnName) {
			this.name = hymnName
		}
	},
	computed: {
		disableUpdateBtn() {
			return (this.position === "" || this.name === "") ? true : false;
		},
		speakerTypeEqHymn() {
			return this.speaker_type === "Hymn" ? true : false;
		}
	},
	watch: {
		item: {
		 	handler(newItem, oldItem) {
				this.position = newItem.position
				this.name = newItem.name
				this.speaker_type = newItem.speaker_type
			},
			immediate: true
		}
	}
}
</script>