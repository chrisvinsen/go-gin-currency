$(function() {
    $('#base_money').text(base_money)

    $('#base_money').on('keyup', function() {
        let money = $(this).text().replace(",", ".")
        money = Number(money)
        if (isNaN(money)) {
            alert("Please enter the correct amount of money")
            base_money = 0
        } else {
            base_money = money
        }

        localStorage.setItem("base_money", base_money)

        $('.card').each(function(index) {
            let rate = Number($(this).find('.rate').text());
            $(this).find('.value').text((rate * base_money).toFixed(2))
        })
    })

    $('#modal-add-currency form').on('submit', function(e) {
        e.preventDefault()

        let index = $('#currency_code').val()
        if (index) {
            addCurrency(index);
            $(`#currency_code option[value="${index}"]`).hide();
        }

        $('#currency_code').prop('selectedIndex',0);
        $('#modal-add-currency').modal('toggle')
    })

    $(document).on('click', '.card .btn-close', function() {
        let index = $(this).data('index')

        currency[index].active = false
        localStorage.setItem("currency", JSON.stringify(currency))

        $(this).closest('.card').remove()
        $(`#currency_code option[value="${index}"]`).show();
    })
})

function addCurrency(index) {
    currency[index].active = true
    localStorage.setItem("currency", JSON.stringify(currency))

    const htmlCurrencyCard = `
        <div class="card col-12">
            <div class="row">
                <div class="col-10 left">
                    <div class="d-flex justify-content-between">
                        <h5 class="code">${currency[index].code}</h5>
                        <h5 class="value">${(currency[index].rate * base_money).toFixed(2)}</h5>
                    </div>
                    <p>${currency[index].code} - ${currency[index].name}</p>
                    <p class="details">1 ${base} = ${currency[index].code} <span class="rate">${currency[index].rate}</span></p>
                </div>
                <div class="col-2 right d-flex align-items-center justify-content-center">
                    <button type="button" class="btn-sm btn-close" data-index="${index}"></button>
                </div>
            </div>
        </div>`
        
    $('.container-btn').before(htmlCurrencyCard)
}
