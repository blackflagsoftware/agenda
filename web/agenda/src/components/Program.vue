<script setup>
import axios from "axios"
import SpeakerType from "./SpeakerType.vue";
import Hymn from "./Hymn.vue";
import { getTransitionRawChildren } from "vue";
</script>

<template>
	<div style="margin-bottom:10px">
		<b>Speakers</b>
	</div>
    <v-checkbox label="Is this a Fast Sunday?" v-model="fastSunday" @change="clickFastSunday"></v-checkbox>
	<div v-if="!fastSunday">
		<v-dialog v-model="modalDialog" max-width="500px">
			<template v-slot:activator="{ props }">
				<v-btn color="blue" dark class="mb-2" v-bind="props">Add Speaker</v-btn>
			</template>
				<v-card>
					<v-card-title>
						<span class="text-h5">{{ modalTitle }}</span>
					</v-card-title>
					<v-card-text>
						<v-container>
							<v-row>
								<div style="width: 120px; margin-right: 20px;">
									<v-select label="Position" variant="outlined" density="compact" v-model="editSpeaker.position" :items="['1', '2', '3', '4', '5', '6', '7', '8', '9', '10']"></v-select>
								</div>
								<SpeakerType :value-in="editSpeaker.speaker_type" @speaker-type-out="speakerTypeOut" /> 
							</v-row>
							<v-row>
								<Hymn :hymns="hymns" :hymn-number-in="editSpeaker.name" @hymn-number-out="hymnNumberOut" v-if="speakerTypeEqHymn" label="Hymn" />
								<v-textarea label="Musical Number" variant="outlined" density="compact" v-model="editSpeaker.name" v-else-if="speakerTypeEqMusical" />
								<v-text-field label="Name" variant="outlined" density="compact" v-model="editSpeaker.name" v-else></v-text-field>
							</v-row>
							<v-card-actions>
								<v-spacer></v-spacer>
								<v-btn color="blue" variant="text" @click="cancelDialog">Cancel</v-btn>
								<v-btn color="blue" variant="text" @click="saveDialog">Save</v-btn>
							</v-card-actions>
						</v-container>
					</v-card-text>
			</v-card>
		</v-dialog>
		<v-sheet v-if="showNoSpeakers">No Speakers</v-sheet>
		<v-form>
			<v-row>
				<v-col cols="12" md="1"><h3>Position</h3></v-col>
				<v-col cols="12" md="1"><h3>Type</h3></v-col>
				<v-col cols="12" md="8"><h3>Name</h3></v-col>
			</v-row>
			<template v-for="(s, idx) in speakers" :key="idx">
				<v-row>
					<v-col cols="12" md="1">{{ s.position }}</v-col>
					<v-col cols="12" md="1">{{ s.speaker_type }}</v-col>
					<v-col cols="12" md="8">{{ staticSpeakerName(s.name, s.speaker_type) }}</v-col>
					<v-col cols="12" md="1"><v-btn @click="updateSpeakerClick(idx)" color="blue">Update</v-btn></v-col>
					<v-col cols="12" md="1"><v-btn @click="deleteSpeakerClick(idx)" color="red">Delete</v-btn></v-col>
				</v-row>
			</template>
			<v-dialog v-model="modalDelete" max-width="530px">
				<v-card>
					<v-card-title class="text-h5">Are you sure you want to delete this speaker?</v-card-title>
					<v-card-actions>
					<v-spacer></v-spacer>
					<v-btn color="blue" variant="text" @click="cancelDeleteClick">Cancel</v-btn>
					<v-btn color="blue" variant="text" @click="deleteDeleteClick">OK</v-btn>
					<v-spacer></v-spacer>
					</v-card-actions>
				</v-card>
			</v-dialog>
		</v-form>
	</div>
</template>

<script>
export default {
	name: "Program",
	data() {
		return {
     	fastSunday: false,
			modalDialog: false,
			modalDelete: false,
			modalTitle: "",
			localAgenda: null,
			position: "",
			name: "",
			speakers: [],
			speakerType: "",
			defaultSpeaker: {
				position: "",
				speaker_type: "",
				name: ""
			},
			editSpeaker: {
				position: "",
				speaker_type: "",
				name: ""
			},
			editIndex: -1
		}
	},
	props: [
		"agenda",
		"hymns"
	],
	emits: {
		refreshAgenda: null,
	},
	mounted() {
		this.fastSunday = this.localAgenda.fast_sunday
		this.getSpeakers()
	},
	methods: {
		getSpeakers: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/speaker/search", {"search": [{"column": "date", "value": this.localAgenda.date, "compare": "="}]})
			.then(response => {
				this.speakers = response.data.data
			})
			.catch(error => {
				console.log(error)
			})
		},
		clickFastSunday: function() {
			const obj = {date: this.localAgenda.date, fast_sunday: this.fastSunday}
			this.$emit("refreshAgenda", obj)
		},
		clickAddSpeaker: function() {
			this.modalTitle = "New Speaker"
			this.modalDialog = true
		},
		updateSpeakerClick: function(idx) {
			this.modalTitle = "Update Speaker"
			this.editIndex = idx
			this.editSpeaker = Object.assign({}, this.speakers[idx])
			this.modalDialog = true
		},
		deleteSpeakerClick: function(idx) {
			this.editIndex = idx
			this.modalDelete = true
		},
		saveDialog: function() {
			const obj = {id: this.editSpeaker.id, date: this.localAgenda.date, position: this.editSpeaker.position, speaker_type: this.editSpeaker.speaker_type, name: this.editSpeaker.name}
			if (this.editIndex === -1) {
				axios.post(import.meta.env.VITE_API_URL + "/v1/speaker", obj)
				.then(() => {
					this.$nextTick(() => {
						this.editSpeaker = Object.assign({}, this.defaultSpeaker)
						this.editIndex = -1
					})
					this.getSpeakers()
					this.modalDialog = false
				})
				.catch(error => {
					console.log(error);
				})
			} else {
				axios.patch(import.meta.env.VITE_API_URL + "/v1/speaker", obj)
				.then(() => {
					this.$nextTick(() => {
						this.editSpeaker = Object.assign({}, this.defaultSpeaker)
						this.editIndex = -1
					})
					this.getSpeakers()
					this.modalDialog = false
				})
				.catch(error => {
					console.log(error);
				})
			}
		},
		cancelDialog: function() {
			this.modalDialog = false
			this.$nextTick(() => {
				this.editSpeaker = Object.assign({}, this.defaultSpeaker)
				this.editIndex = -1
			})
		},
		cancelDeleteClick: function() {
			this.modalDelete = false
			this.$nextTick(() => {
				this.editSpeaker = Object.assign({}, this.defaultSpeaker)
				this.editIndex = -1
			})
		},
		deleteDeleteClick: function() {
			axios.delete(import.meta.env.VITE_API_URL + "/v1/speaker/" + this.speakers[this.editIndex].id)
			.then(() => {
				this.speakers.splice(this.editIndex, 1)
				this.cancelDeleteClick();
			})
			.catch(error => {
				console.log(error);
			})
		},
		speakerTypeOut: function(speakerType) {
			this.editSpeaker.speaker_type = speakerType
			this.editSpeaker.name = ""
		},
		hymnNumberOut: function(hymnName) {
			this.editSpeaker.name = hymnName
		},
		staticSpeakerName: function(name, speakerType) {
			var hymnName = ""
			if (speakerType === "Hymn") {
				hymnName = this.hymns[Number(name)-1].name
				return hymnName
			}
			return name
		}
	},
	computed: {
		disableSaveBtn() {
			return (this.position === "" || this.speakerTypeOut === "" || this.name === "") ? true : false;
		},
		showNoSpeakers() {
			return this.speakers.length === 0 ? true : false;
		},
		speakerTypeEqHymn() {
			return this.editSpeaker.speaker_type === "Hymn" ? true : false;
		},
		speakerTypeEqMusical() {
			return this.editSpeaker.speaker_type === "Musical Number" ? true : false;
		}
	},
	watch: {
		agenda: {
			handler(newAgenda, oldAgenda) {
				this.localAgenda = newAgenda;
				this.fastSunday = this.localAgenda.fast_sunday;
				this.getSpeakers();
			},
			immediate: true
		}
	}
}
</script>