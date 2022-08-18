<template>
  <v-container>
    <v-row class="form-container">
      <v-col cols="12" md="6">
        <v-text-field
          append-icon="mdi-magnify"
          v-model="search"
          :label="$t('search')"
          outlined
          clearable
        ></v-text-field>
        <v-skeleton-loader
          class="mb-5"
          type="article: heading, paragraph"
          v-if="loading"
        ></v-skeleton-loader>
        <v-skeleton-loader
          type="article: heading, paragraph"
          v-if="loading"
        ></v-skeleton-loader>
        <v-treeview v-else :search="search" item-key="name" :items="items">
          <template v-slot:label="{ item, open }">
            <div class="flex-space">
              <div class="control-text">{{ item.name }}</div>
              <div>
                <v-icon @click="updateCategory(item)">
                  {{
                    open ? "mdi-folder-edit-outline" : "mdi-folder-edit-outline"
                  }}
                </v-icon>
                <v-icon
                  @click="selectCategory(item)"
                  v-if="item.level < 2"
                  class="ma-2"
                >
                  {{ open ? "mdi-plus-circle" : "mdi-plus-circle" }}
                </v-icon>
                <v-icon @click="deleteCategory(item)">
                  {{ open ? "mdi-delete" : "mdi-delete" }}
                </v-icon>
              </div>
            </div>
          </template>
        </v-treeview>
      </v-col>
      <v-col cols="12" md="6">
        <v-form value ref="form" class="mainForm">
          <v-row>
            <v-col cols="12">
              <template>
                <v-text-field
                  ref="focusEnglishName"
                  outlined
                  :label="$t('name')"
                  v-model="form.groupNameEn"
                  required
                  clearable
                  :rules="inputRules"
                ></v-text-field>
              </template>
            </v-col>
            <v-col cols="12">
              <template>
                <v-text-field
                  outlined
                  ref="focusArabicName"
                  :label="$t('arabicName')"
                  v-model="form.groupName"
                  required
                  clearable
                  :rules="inputRules"
                ></v-text-field>
              </template>
            </v-col>
            <v-col cols="12">
              <template>
                <v-text-field
                  outlined
                  :label="$t('code')"
                  readonly
                  v-model="form.code"
                ></v-text-field>
              </template>
            </v-col>
            <v-col cols="12">
              <v-combobox
                v-if="disable"
                outlined
                :disabled="disable"
                clearable
                label="data Not found"
              ></v-combobox>
              <v-combobox
                v-else
                outlined
                clearable
                v-model="selectedCategory"
                :disabled="disable"
                :items="comboItems"
                item-text="name"
                @change="comboChanged"
                :label="$t('parentCode')"
              ></v-combobox>
            </v-col>
            <v-col cols="12">
              <template>
                <v-select
                  outlined
                  :items="staticData"
                  v-model="form.type"
                  :label="$t('categoryType')"
                  dense
                  clearable
                ></v-select>
              </template>
            </v-col>
            <v-col cols="12">
              <v-file-input
                outlined
                accept="image/*"
                v-model="form.image"
                :label="$t('image')"
              ></v-file-input>
            </v-col>
            <v-col cols="12">
              <v-btn depressed width="100%" color="primary" @click="submitData">
                {{ $t("submit") }}
              </v-btn>
            </v-col>
            <v-row>
              <template>
                <div class="text-center">
                  <v-snackbar
                    dark
                    top
                    right
                    :color="snackColor"
                    v-model="snackbar"
                    :timeout="timeout"
                  >
                    {{ text }}
                  </v-snackbar>
                </div>
              </template>
            </v-row>
          </v-row>
        </v-form>
      </v-col>
      <template>
        <v-row justify="center">
          <v-dialog v-model="dialog" persistent max-width="290">
            <v-card>
              <v-card-title class="text-h4"> {{ $t("Warning") }} </v-card-title>
              <v-card-text>{{ $t("warningMsg") }}</v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="green darken-1" text @click="dialog = false">
                  {{ $t("close") }}
                </v-btn>
                <v-btn color="green darken-1" text @click="dialog = false">
                  {{ $t("agree") }}
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-row>
      </template>
    </v-row>
  </v-container>
</template>

<script>
import Vue from "vue";
import axios from "axios";
import { generateColors, switchMode } from "@/utils/helpers";
import groupsData from "../../utils/groups/groups";
import { comboType, itemType } from "@/utils/groups/comboInterface";
const groups = new groupsData();

export default Vue.extend({
  name: "HomeView",
  data() {
    return {
      form: {
        groupNameEn: "",
        groupName: "",
        // image: null,
        parentCode: 0,
        code: "",
        // type: "",
      },
      search: "",
      disable: true,
      postRequestSuccess: true,
      dialog: false,
      selectedCategory: "",
      snackbar: true,
      updateGroupCode: 0,
      loading: true,
      failedPost: "",
      snackColor: "",
      snackbar: false,
      text: "",
      timeout: 2000,
      value: "add the category",
      inputRules: [(v) => !!v || "Item is required"],
      comboItems: groups.comboItem,
      staticData: ["خامات", "منتج نهائي"],
      items: [],
    };
  },
  methods: {
    fillForm(item) {
      this.updateGroupCode = item.groupCode;
      this.form.groupName = item.groupName;
      this.form.groupNameEn = item.groupNameEn;
      this.form.code = item.code;
      this.form.parentCode = item.parentCode;
    },
    updateCategory(item) {
      this.resetUpdateForm();
      this.fillForm(item);
      console.log(item.parentCode);
      const category = this.comboItems.filter((comboitem) => {
        return comboitem.code == item.parentCode;
      });
      this.selectedCategory = category.length > 0 ? category[0].name : "";
    },
    // updateRequest(groupCode) {
    //   axios
    //     .put(`http://192.168.1.40:5000/api/group/${groupCode}`, this.form)
    //     .then((res) => {
    //       console.log(res);
    //     });
    // },
    async Request() {
      const url =
        this.updateGroupCode == 0
          ? `http://192.168.1.40:5000/api/group`
          : `http://192.168.1.40:5000/api/group/${this.updateGroupCode}`;
      const method = this.updateGroupCode == 0 ? "post" : "put";
      await axios[method](url, this.form)
        .then((res) => {
          if (this.updateGroupCode == 0) {
            this.text = "category added successfully";
            this.snackColor = "#28a745";
          } else {
            this.text = "category updated successfully";
            this.snackColor = "#28a745";
          }
          console.log(res);
        })
        .catch((err) => {
          this.text = err.message;
          this.postRequestSuccess = false;
          console.log(err.message);
        });
      this.snackbar = !this.snackbar;
      this.$refs.form.reset();
      this.$refs.form.resetValidation();
      this.getData();
      this.comboChanged(null);
    },
    submitData() {
      if (this.$refs.form.validate()) {
        this.Request();
      } else {
        setTimeout(() => {
          this.$refs.form.resetValidation();
        }, 5000);
      }
    },
    convertCurrentItemToStr(item) {
      let strItem = item.toString();
      if (item <= 9) {
        strItem = `0${item}`;
      }
      return strItem;
    },
    generateGroupCode(currentItem) {
      return `${currentItem.code}${this.convertCurrentItemToStr(
        currentItem.childrenLength + 1
      )}`;
    },
    comboChanged(item) {
      this.updateGroupCode = 0;
      if (item == null || typeof item != "object") {
        this.form.code =
          this.items.length + 1 > 9
            ? `${this.items.length + 1}`
            : `0${this.items.length + 1}`;
        this.form.parentCode = "";
        this.resetAddForm();
        return;
      }

      if (item.location[1] == -1) {
        this.generateCode(this.items[item.location[0]]);
      } else {
        this.generateCode(
          this.items[item.location[0]].children[item.location[1]]
        );
      }
    },
    resetAddForm() {
      // await this.$refs.form.reset();
      this.$refs.form.resetValidation();
      this.$refs.focusEnglishName.focus();
      this.$refs.focusEnglishName.reset();
      this.$refs.focusArabicName.reset();
    },
    resetUpdateForm() {
      this.$refs.form.resetValidation();
      this.$refs.focusEnglishName.focus();
    },
    selectCategory(item) {
      this.updateGroupCode = 0;
      this.resetAddForm();
      this.selectedCategory = item.groupNameEn;
      this.generateCode(item);
    },
    deleteCategory(item) {
      this.dialog = !this.dialog;
    },
    generateCode(item) {
      this.form.code = this.generateGroupCode(item);
      console.log(item.code);
      this.form.parentCode = item.code;
    },
    async getData() {
      const req = await axios.get(
        `http://192.168.1.40:5000/api/group/hierarchy?lang=${localStorage.getItem(
          "locale"
        )}`
      );
      if (req.data == null) {
        this.items = [];
      } else {
        this.items = req.data;
      }
      if (this.loading != false && this.disable != false) {
        this.loading = !this.loading;
        this.disable = !this.disable;
      }
    },
  },
  async mounted() {
    await this.getData();
    console.log(this.items);
    groups.insertComboData(this.items);
    this.comboChanged(null);
  },
});
</script>
//
