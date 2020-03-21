package {{packageName}}.entity.{{moduleName}};

import com.zjzhd.entity.BaseEntity;
import javax.persistence.*;
{% if hasBigDecimal  %}
import java.math.BigDecimal;{% endif %}{% if hasDate  %}import java.util.Date;{% endif %}{% if hasTime  %}import java.sql.Time;{% endif %}

/**
 * Description:{{ table.TableComment }}
 * <p>
 *
 * @author {{ authorName }}
 * @date {{ genTime }}
 */
@Entity
@Table(name = "{{ table.TableName }}")
public class {{ className }} extends BaseEntity {

    {% for column in listColumns %}{% if !column.IsJoinColumn%}
    /**
     * {{ column.ColumnComment }}
     */
    {% if column.ColumnKey == "PRI" %}
    @Id
    @Column(name = "id", nullable = false)
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    {% else %}
    @Column(name = "{{ column.ColumnName }}"{% if column.Length > 0 %}, length = {{ column.Length }}{% endif%})
    {% endif %}
    private {{ column.JavaType }} {{ column.FieldName}};
    {% endif %}
    {% endfor %}


    {% for column in listColumns %}{% if !column.IsJoinColumn%}
    public {{ column.JavaType }} get{{ column.Uppercase1th }}() {
        return {{ column.FieldName}};
    }

    public void set{{ column.Uppercase1th}}({{ column.JavaType }} {{ column.FieldName}}) {
        this.{{ column.FieldName}} = {{ column.FieldName}};
    }
    {% endif %}{% endfor %}



}
