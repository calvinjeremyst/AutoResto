<template>
  <div
    class="columns is-centered"
    style="align-items: center; min-height: 100vh"
  >
    <div class="column is-four-fifths">
      <div class="card">
        <div class="card-content">
          <h1 class="title cart-header" style="color: red">
            <b>Menu Utama</b>
          </h1>

          <div class="content" style="margin-bottom: 0">
            <b-table :data="data ? data : []" :hoverable="true">
              <b-table-column
                v-slot="props"
                field="id"
                label="No"
                numeric
                width="40"
              >
                {{ props.index + 1 }}
              </b-table-column>

              <b-table-column v-slot="props" field="Name" label="Name">
                {{ props.row.nama }}
              </b-table-column>

              <b-table-column
                v-slot="props"
                field="description"
                label="Description"
                width="200"
              >
                {{ props.row.deskripsi }}
              </b-table-column>

              <b-table-column v-slot="props" field="image" label="Image">
                <figure style="text-align: center; margin: 0">
                  <img
                    :src="`images/menu/${props.row.gambar}`"
                    class="image-product"
                    width="200px"
                  />
                </figure>
              </b-table-column>

              <b-table-column
                v-slot="props"
                field="quantity"
                label="Quantity"
                width="125"
              >
                <b-field>
                  <b-numberinput
                    v-model="props.row.quantity"
                    controls-position="compact"
                    min="1"
                    size="is-small"
                  >
                  </b-numberinput>
                </b-field>
              </b-table-column>

              <b-table-column v-slot="props" field="price" label="Price">
                Rp. {{ parseInt(props.row.harga) }}
              </b-table-column>

              <b-table-column v-slot="props" field="total" label="Total"
                ><span v-if="typeof props.row.harga !== 'undefined'">
                  Rp. {{ parseInt(props.row.harga * props.row.quantity) }}
                </span>
                <span v-else> - </span>
              </b-table-column>

              <b-table-column v-slot="props" field="action" label="Action">
                <b-button
                  icon-right="trash"
                  type="is-danger"
                  @click="propsDeleteItem(props.row)"
                />
              </b-table-column>

              <template slot="empty">
                <section class="section">
                  <div class="content has-text-grey has-text-centered">
                    <p>Cart is empty!</p>
                  </div>
                </section>
              </template>
            </b-table>
          </div>

          <hr style="background-color: black; margin: 24px 0" />

          <div class="level is-flex" style="padding: 0 24px">
            <p><strong>Total All Item</strong></p>
            <p>
              <strong>Rp {{ parseInt(calculateTotal()) }}</strong>
            </p>
          </div>

          <hr style="background-color: black; margin: 24px 0" />

          <div class="level">
            <div class="level-right is-flex" style="width: 100%">
              <b-button
                style="margin-right: 12px"
                tag="router-link"
                to="/menu"
                type="is-light"
                ><b>Add another Pizza</b></b-button
              >
              <b-button
                style="width: 180px"
                type="is-primary"
                ><b>Masuk</b></b-button
              >
            </div>
          </div>
        </div>
      </div>
    </div>

    <b-modal
      v-model="isComponentModalActive"
      :destroy-on-hide="false"
      has-modal-card
      trap-focus
    >
      <form action="">
        <div class="modal-card" style="width: 500px">
          <header class="modal-card-head">
            <p class="modal-card-title">Details</p>
          </header>
          <section class="modal-card-body">
            <b-field label="Total">
              <b-input :value="getTotalBill" disabled name="form.total">
              </b-input>
            </b-field>
            <b-field label="email">
              <b-input type="email" maxlength="30"> </b-input>
            </b-field>
            <b-field label="nama">
              <b-input v-model="nama"></b-input>
            </b-field>
            <b-field label="alamat">
              <b-input v-model="nama"></b-input>
            </b-field>
          </section>
          <footer class="modal-card-foot">
            <b-button
              style="color: white"
              type="is-secondary"
              @click="isComponentModalActive = false"
              >Cancel</b-button
            >
            <b-button
              type="is-primary"
              @click="postTransaction()"
              >Continue to Pay
            </b-button>
          </footer>
        </div>
      </form>
    </b-modal>
  </div>
</template>

<script>
import CartService from "@/services/CartService";
import axios from "axios";

export default {
  data() {
    const data = [];
    return {
      form: {
        method: "",
        total: 1000000,
      },

      cartService: new CartService(),

      // model for table
      quantities: [],

      // Table
      data,

      // modal
      isComponentModalActive: false,
    };
  },
  mounted() {
    this.fetchData();
  },
  computed: {
    getTotalBill: function () {
      return "Rp " + parseInt(this.calculateTotal());
    },
  },
  methods: {
    calculateTotal() {
      var total = 0;
      this.data.forEach(function (item) {
        if (typeof item.quantity !== "undefined") {
          total += item.quantity * item.harga;
        }
      });
      return total;
    },
    checkValidity() {
      var isValid = true;
      this.data.forEach(function (item) {
        if (typeof item.quantity === "undefined") {
          isValid = false;
          return;
        }
      });
      return isValid;
    },
    async fetchData() {
      this.data = this.cartService.getCurrentCart();
    },
    propsDeleteItem(item) {
      this.$buefy.dialog.confirm({
        title: "Delete this item from cart?",
        message:
          "Item that has been <b>deleted</b> will not be count as your item again!",
        type: "is-danger",
        hasIcon: true,
        icon: "times-circle",
        iconPack: "fa",
        cancelText: "No",
        confirmText: "Yes",
        onConfirm: () => (item.quantity = 0),
      });
    },
    async postTransaction() {
      const orders = this.cartService
        .getCurrentCart()
        .map(({ id, quantity }) => ({
          pizza_id: id,
          quantity,
        }));
      console.log(orders);
      await axios.post("/order/create", orders);
      alert("Payment Successful!");
      this.isComponentModalActive = false;
    },
  },
};
</script>

<style scoped>
.column {
  padding-top: 24px;
}

.cart-header {
  border-bottom: 1px solid black;
  padding-bottom: 24px;
}

button {
  background-image: linear-gradient(to right, darkslategray, gray);
}
</style>
